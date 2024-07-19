package tui

import (
	"fmt"
	"os"

	"togo/task"
	"togo/tui/constants"

	tea "github.com/charmbracelet/bubbletea"
)

func Start(tr task.GormRepository) {
	// Set the global variable for our data repository
	constants.Tr = &tr

	// Initialize task view and start program
	m := InitTask()
	constants.Program = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.Program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
