package main

import (
	"ds-protector/internal/views"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_, err := tea.NewProgram(views.NewApp(), tea.WithAltScreen()).Run()
	if err != nil {
		return
	}
}
