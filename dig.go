package main

import (
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// runDigCommand runs the 'dig' command asynchronously and returns the result
func runDigCommand(domain string) tea.Cmd {
	return func() tea.Msg {
		out, err := exec.Command("dig", "+short", domain).Output()
		if err != nil {
			return digResultMsg{output: "", err: err}
		}
		return digResultMsg{output: strings.TrimSpace(string(out)), err: nil}
	}
}
