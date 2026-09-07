#!/usr/bin/env python3
"""Where does a player stand among everyone else at their rank?

Reads the AGA TDList - the rating list every tournament director already
downloads - and reports a player's position inside their own rank band.

    python3 aga-rank.py tdlist.tsv --fetch          # download today's list
    python3 aga-rank.py tdlist.tsv "Steinberg"      # by name (substring)
    python3 aga-rank.py tdlist.tsv 17531            # by AGA ID
    python3 aga-rank.py tdlist.tsv 17531 --since 2024   # only recently-rated players

Public domain / CC0.  Python 3.9+, standard library only.
"""

import argparse
import csv
import math
import re
import sys
import urllib.request

# The AGA's stable handle for the tab-separated "A" variant of the list.  It is
# NOT an HTTP redirect: usgo.org serves a page carrying a JavaScript redirect to
# an Azure host whose name changes, so resolve it at runtime rather than pasting
# the Azure URL anywhere.  That host also rejects a default urllib User-Agent.
HANDLE = "https://www.usgo.org/TDListA"
UA = ("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 "
      "(KHTML, like Gecko) Chrome/126.0 Safari/537.36")


def get(url, timeout=90):
    req = urllib.request.Request(url, headers={"User-Agent": UA})
    return urllib.request.urlopen(req, timeout=timeout).read()


def fetch(path):
    page = get(HANDLE, timeout=30).decode("utf-8", "replace")
    m = re.search(r"""https://[^\s'"]*GenerateTDList[A-Z]?""", page)
    if not m:
        sys.exit("no GenerateTDList URL found on " + HANDLE)
    data = get(m.group(0))
    if len(data) < 100_000:          # a real list is ~1 MB; smaller means error page
        sys.exit("suspiciously small download (%d bytes)" % len(data))
    open(path, "wb").write(data)
    print("wrote %s (%.1f MB) from %s" % (path, len(data) / 1e6, m.group(0)))


def rank_of(rating):
    """AGA rating -> rank string.  Truncates toward zero: 4.99 is a 4d, not a 5d.
    The scale has no 0: ratings run ... -2 (2k) -1 (1k) | 1 (1d) 2 (2d) ..."""
    n = math.trunc(rating)
    return "%dd" % n if rating > 0 else "%dk" % -n


def load(path):
    """Every row of the TDList that carries a rating.

    Columns: name, AGA id, membership type, rating, rating date, chapter,
    state, sigma, member-since.  A blank rating means the player has never
    been rated - those rows are not players in this sense, so drop them."""
    out = []
    for row in csv.reader(open(path, encoding="utf-8", errors="replace"), delimiter="\t"):
        if len(row) < 5 or not row[3].strip():
            continue
        try:
            rating = float(row[3])
        except ValueError:
            continue
        year = int(row[4].split("/")[-1]) if "/" in row[4] else 0
        out.append({"name": row[0], "id": row[1], "rating": rating, "year": year})
    return out


def report(players, who, since=0):
    pool = [p for p in players if p["year"] >= since]
    hits = [p for p in players
            if p["id"] == who or who.lower() in p["name"].lower()]
    if not hits:
        sys.exit("no player matching %r" % who)
    for p in hits:
        band = rank_of(p["rating"])
        peers = sorted(q["rating"] for q in pool if rank_of(q["rating"]) == band)
        if p["rating"] not in peers:      # filtered out by --since
            print("%s (%s): %s, excluded by --since %d" % (p["name"], p["id"], band, since))
            continue
        pos = peers.index(p["rating"]) + 1               # 1 = weakest in the band
        weaker = sum(1 for q in pool if q["rating"] < p["rating"])
        print("%s (%s)  %.5f  %s" % (p["name"], p["id"], p["rating"], band))
        print("   #%d of %d %s players from the bottom (#%d from the top)"
              % (pos, len(peers), band, len(peers) - pos + 1))
        print("   stronger than %d of %d rated players (%.1f%%)"
              % (weaker, len(pool), 100.0 * weaker / len(pool)))


def main():
    ap = argparse.ArgumentParser(description=__doc__,
                                 formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("tdlist", help="path to a TDList .tsv snapshot")
    ap.add_argument("player", nargs="?", help="AGA id, or part of a name")
    ap.add_argument("--fetch", action="store_true", help="download a fresh list to TDLIST first")
    ap.add_argument("--since", type=int, default=0, metavar="YEAR",
                    help="only count players rated in YEAR or later")
    a = ap.parse_args()
    if a.fetch:
        fetch(a.tdlist)
    if a.player:
        report(load(a.tdlist), a.player, a.since)


if __name__ == "__main__":
    main()
