package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/keybase/go-keychain"
)

// Allowed callers - matched against process cmdline (substring match)
// These run via nix-shell shebang which strips absolute paths,
// so we match on the relative form that appears in `ps` output.
var allowedScripts = []string{
	"bin/deploy",
	"bin/connect",
	"bin/diff",
	"bin/pull",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: wmgc-secret <get|check>")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "get":
		if !isCallerAllowed() {
			fmt.Fprintln(os.Stderr, "❌ Caller not in allowed list")
			fmt.Fprintln(os.Stderr, "Process tree:")
			dumpProcessTree(os.Stderr)
			fmt.Fprintf(os.Stderr, "Allowed scripts: %v\n", allowedScripts)
			os.Exit(1)
		}
		secret, err := getSecret()
		if err != nil {
			fmt.Fprintln(os.Stderr, "❌ Keychain error:", err)
			os.Exit(1)
		}
		fmt.Print(secret) // No newline - for shell substitution

	case "check":
		if !isCallerAllowed() {
			fmt.Fprintln(os.Stderr, "❌ Caller not in allowed list")
			fmt.Fprintln(os.Stderr, "Process tree:")
			dumpProcessTree(os.Stderr)
			fmt.Fprintf(os.Stderr, "Allowed scripts: %v\n", allowedScripts)
			os.Exit(1)
		}
		_, err := getSecret()
		if err != nil {
			fmt.Println("❌ Keychain access failed:", err)
			os.Exit(1)
		}
		fmt.Println("✅ Keychain access OK")

	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", cmd)
		os.Exit(1)
	}
}

// dumpProcessTree prints the process tree to the given writer for debugging
func dumpProcessTree(w *os.File) {
	pid := os.Getppid()
	for i := 0; i < 10 && pid > 1; i++ {
		cmdline := getProcessCmdline(pid)
		if cmdline == "" {
			break
		}
		fmt.Fprintf(w, "  PID %d: %s\n", pid, cmdline)
		pid = getParentPID(pid)
	}
}

// isCallerAllowed walks up the process tree looking for an allowed script
func isCallerAllowed() bool {
	pid := os.Getppid()

	// Walk up to 10 levels (plenty for any reasonable call chain)
	for i := 0; i < 10 && pid > 1; i++ {
		cmdline := getProcessCmdline(pid)
		for _, script := range allowedScripts {
			if strings.Contains(cmdline, script) {
				return true
			}
		}
		pid = getParentPID(pid)
	}
	return false
}

// getProcessCmdline gets the command line of a process via ps
func getProcessCmdline(pid int) string {
	out, err := exec.Command("ps", "-o", "args=", "-p", strconv.Itoa(pid)).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// getParentPID gets the parent PID of a process
func getParentPID(pid int) int {
	out, err := exec.Command("ps", "-o", "ppid=", "-p", strconv.Itoa(pid)).Output()
	if err != nil {
		return 0
	}
	ppid, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 0
	}
	return ppid
}

// getSecret retrieves the FTP password from keychain
func getSecret() (string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService("wmgc.massgo.org")
	query.SetAccount("wmgcadmin")
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)

	results, err := keychain.QueryItem(query)
	if err != nil {
		return "", err
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no keychain item found")
	}
	return string(results[0].Data), nil
}
