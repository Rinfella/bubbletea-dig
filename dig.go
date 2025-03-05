package main

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// runDigCommand runs the 'dig' command asynchronously and returns the result
func runDigCommand(domain string) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("dig", "+short", domain)
		output, err := cmd.CombinedOutput()
		return digResultMsg{
			output: string(output),
			err:    err,
		}
	}
}
