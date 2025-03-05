package main

import (
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
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			return m, runDigCommand(m.input.Value())
		}

	case digResultMsg:
		if msg.err != nil {
			m.err = msg.err
			m.result = ""
		} else {
			m.result = msg.output
			m.err = nil
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}
