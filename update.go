package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// digResultMsg represents the output of the dig command
type digResultMsg struct {
	output string
	err    error
}

// Update function: handles user input, executes commands and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.input.Value() != "" {
				m.loading = true
				m.result = ""
				return m, runDigCommand(m.input.Value()) // Call dig
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case digResultMsg:
		m.loading = false
		if msg.err != nil {
			m.result = "Error: " + msg.err.Error()
		} else {
			m.result = strings.TrimSpace(msg.output)
		}
	}

	// Handle text input updates
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
