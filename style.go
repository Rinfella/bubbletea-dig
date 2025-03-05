package main

import "github.com/charmbracelet/lipgloss"

// Color Palette
const (
	colorPrimary   = "#FFD700" // Gold
	colorSecondary = "#87CEFA" // Light Blue
	colorError     = "#FF5555" // Red
	colorInput     = "#00FF00" // Green
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPrimary)).
			Bold(true).
			Underline(true).
			Padding(0, 1)

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorInput)).
			Padding(0, 1)

	resultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorSecondary)).
			Italic(true).
			Padding(1, 2)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorError)).
			Bold(true)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			BorderForeground(lipgloss.Color(colorPrimary))
)
