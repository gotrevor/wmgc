package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/keybase/go-keychain"
)

// Allowed callers - scripts that may request the secret
var allowedScripts = []string{
	"/Users/gotrevor/src/wmgc/bin/connect",
	"/Users/gotrevor/src/wmgc/bin/deploy",
	"/Users/gotrevor/src/wmgc/bin/diff",
	"/Users/gotrevor/src/wmgc/bin/pull",
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
			os.Exit(1) // Silent rejection
		}
		secret, err := getSecret()
		if err != nil {
			fmt.Fprintln(os.Stderr, "keychain error:", err)
			os.Exit(1)
		}
		fmt.Print(secret) // No newline - for shell substitution

	case "check":
		if !isCallerAllowed() {
			fmt.Println("❌ Caller not in allowed list")
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
