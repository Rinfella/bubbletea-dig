package main

import "fmt"

func (m model) View() string {
	s := "Enter a domain & press Enter:\n" + m.input.View() + "\n\n"

	if m.loading {
		s += "Looking up domain...\n"
	} else if m.result != "" {
		s += fmt.Sprintf("Result: \n%s\n", m.result)
	}

	s += "\nPress 'q' to quit."
	return s
}
