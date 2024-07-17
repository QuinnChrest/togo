package tui

import (
	"fmt"
	"os"

	constants "togo/tui/constants"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	m := InitTask()
	constants.Program = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.Program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
