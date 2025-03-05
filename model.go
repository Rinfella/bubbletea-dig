package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input   textinput.Model
	result  string
	loading bool
	err     error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter domain: (e.g., example.com)"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30

	return model{
		input:   ti,
		result:  "",
		loading: false,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
