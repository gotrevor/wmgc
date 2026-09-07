# Pairing without ranks

*Paul Matthews' AccelRat, the Go tournament system that never rounds you off*

[ Claude wrote all of this.  I think it's good stuff - but, treat with care!  -Trevor ]

---

Every so often somebody points out that discrete ranks are an odd thing to build a tournament on.  A player goes 4-1, his AGA rating drops from 8.03 to 7.99, and he "loses a stone" - the games said he is a shade weaker than he was, and the ranking said he crossed a wall.  The bands decide who he plays, what handicap he takes, and which prize he is eligible for, and the wall between two of them fell in an arbitrary place.

The usual reply is that ranks are a practical necessity: pairing, handicaps and prizes all need buckets, so somebody has to round.  That reply is wrong, and there is a working counterexample that has been running for thirty years.

## AccelRat

**Paul G. Matthews** built a Go tournament system that pairs on ratings and never converts them to ranks.  It has no McMahon bar, no bands, no groups you enter at registration.  The old name is **AccelRat** - "Accelerated Ratings" - and it lives today at [accelrat.com](http://www.accelrat.com/) under the name **GTR, Go Tourney Ratings**, still free and still online.

Matthews is not a hobbyist who wrote a pairing script.  He has a Ph.D. in cognitive science and statistics from Stanford and spent nearly thirty years as a research scientist at Bell Labs and Telcordia.  He was **the AGA's ratings statistician for two decades**, and he is the person who redesigned the AGA's core rating algorithm in 1988-89.  He directed the New Jersey Open at Princeton for more than twenty years and a U.S. Go Congress at Rutgers.  The system is what an AGA ratings statistician built when he got to run tournaments the way he thought they should be run.

## How it works

**Ratings are the object, not a source for ranks.**  A rating is a decimal whose integer part is the rank and whose sign is dan or kyu.  There is no zero, so a dan-versus-kyu difference takes a correction: the gap between 1.5 and -1.5 is (1.5 - 1) - (-1.5 + 1) = 1.0, one stone, as it should be.  An *n*-stone handicap is defined to be an *n*-point rating difference, which pins the whole scale up to an additive constant.

**Re-rate every round, then pair top-down.**  After each round the field is re-rated - a Bayesian update where the prior is your previous rating with a spread (sigma) that widens the longer it has been since you played, and widens further if you were never rated.  Then the highest-rated player takes the nearest eligible opponent, and so on down the list.  When somebody runs out of legal opponents the algorithm backtracks and re-chooses for a player already paired; the avoidance and no-repeat constraints relax only when there is provably no solution.

**The "accelerated" part is look-ahead.**  You tell it how many rounds remain, and it pairs all of them in reverse - last round first, working backward - then uses the resulting round as the one you actually play.  The effect is to hold the top rivalries back for the late rounds rather than burning them in round one.

**Handicaps come from a target winning percentage,** not from subtracting ranks.  You pick a number between 50 (full compensation) and 0 (never); around 10 means a handicap appears only across a large gap.  The calculation accounts for how uncertain the two ratings are, which is not the same thing as looking at the difference between them.

**Prizes are clustered afterwards, not declared beforehand.**  This is the part that answers the practical objection, and it is the cleverest idea in the system.  Because there are no bands, award groups are *derived from how the tournament actually played out*.  Start with every player alone, ordered by rating.  Define a group's **coherence** as the fraction of its members' games that were played inside the group.  Repeatedly merge the adjacent pair whose merger has the highest coherence, subject to your minimum and maximum group sizes, until no merge improves things.  The groups that fall out are the sets of players who really did compete with each other.  Within a group, place by wins, and break ties on a **genesis rating**: re-rate everybody from a common seed so the tiebreak reflects only who beat whom, not who showed up with the bigger number.

So the buckets still exist - you cannot hand out four trophies without them - but they are computed from the games instead of drawn on the entry form, and no player's tournament is shaped by which side of a line their rating landed on in March.

## Has anyone actually used it?

Yes, and more recently than you would guess.  The GTR system's public tourney list holds **around 500 events**, and the most recent one is dated **June 13, 2026**.  Checking the years named in the tourney titles: roughly 13 events in 2026, 35 in 2025, 25 in 2024, 51 in 2023, 81 in 2022.

Most of that volume is Feng Yun 9p's Saturday class in New Jersey, which Matthews rates every week - he assists at her school.  But it is not only him.  The **San Francisco Go Club** ran its 2022 and 2023 tournament series on it (Cherry Blossom, Obon, Shinji Dote Memorial, ING Foundation, Mid-Autumn, APEC) and its BadukPop Tournament in 2025 and again in 2026, and a Chicago organizer has been running youth tournaments and classroom practice through it into 2026.

That is a real answer to the "it would never work in practice" objection.  It has worked in practice, at real weekend tournaments with prizes, for a long time.

## Why it isn't how the AGA runs

Two reasons, neither of them "the idea is broken."

The first is rules.  The AGA's tournament rules do not permit ratings-based pairing outright; the commentary to the rules says it in so many words - *"We can give waivers to permit something like Paul Mathews's Accelerated pairing system, which uses ratings instead of ranks.  It is probably premature to switch to ratings for all tournament pairings."*  So it runs by exception, one waiver at a time.

The second is plumbing.  AccelRat no longer interoperates with the current AGA ratings software, so a director who uses it has extra work to get games submitted for rating.  Meanwhile [OpenGotha](https://www.opengotha.info/) became the default across the West - it speaks the AGA's file formats, and it is what almost every North American tournament, including ours, is paired with.  Being the thing that fits the pipeline beat being the thing with the better idea, which is not an unusual ending.

## The point

Discrete ranks are a presentation choice that hardened into an architecture.  The arithmetic underneath them is already continuous - the AGA has rated on a decimal scale since Matthews rebuilt the algorithm in 1989, and your rank is that decimal with the fractional part thrown away.  Pairing, handicapping and prize-giving can all be done from the number itself; somebody built the whole tournament stack that way and ran hundreds of events on it.

The wall a player crosses at 8.00 is not a fact about Go.  It is a fact about a display format we kept.

---

**Sources.**  Documentation and the live application: [accelrat.com](http://www.accelrat.com/) - see its *Background*, *Ratings*, *Pairings* and *Standings* pages; the pre-2005 algorithm pages (`PairAlg.htm`, `Prior.htm`) survive only in the [Internet Archive](https://web.archive.org/web/20101009034817/http://www.accelrat.com/PairAlg.htm).  Matthews' own summary of the rating system he built for the AGA, *Inside the AGA Rating System*, is [mirrored by the French Go Federation](https://ffg.jeudego.org/echelle/aga-rating.txt).  The waiver language is in the [AGA Tournament Rules](https://www.cs.cmu.edu/~wjh/go/rules/tournrules.html), section VI and its commentary.  Usage figures were read off the GTR public tourney list on 2026-09-07.

[Western Massachusetts Go Club](index.html)
