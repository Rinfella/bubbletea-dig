package main

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	title := titleStyle.Render("🔍 DIG TUI")

	input := inputStyle.Render(m.input.View())

	var result string
	if m.err != nil {
		result = errorStyle.Render("❌ Error: " + m.err.Error())
	} else {
		result = resultStyle.Render("📜 Result:\n" + m.result)
	}

	ui := lipgloss.JoinVertical(lipgloss.Left, title, input, result)

	return boxStyle.Render(ui)
}
