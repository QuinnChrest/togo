package tui

import (
	"fmt"
	"os"

	"togo/task"
	"togo/tui/constants"

	tea "github.com/charmbracelet/bubbletea"
)

func Start(tr task.GormRepository) {
	constants.Tr = &tr
	m := InitTask()
	constants.Program = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.Program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
