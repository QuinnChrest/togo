package constants

import (
	"togo/task"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Lipgloss styles for rendering
var (
	HeaderTextStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("7")).
		BorderForeground(lipgloss.Color("#ec42ff")).
		Border(lipgloss.RoundedBorder())
	
	HeaderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ec42ff"))

	ItemStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("7")).
		BorderLeft(true).
		PaddingLeft(2)

	SelectedItemStyle = ItemStyle.
		Foreground(lipgloss.Color("#ec42ff")).
		BorderForeground(lipgloss.Color("#ec42ff")).
		Bold(true)

	ActivePageDotStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ec42ff"))

	InactivePageDotStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"})
)

// Global variables for application
var (
	WindowSize tea.WindowSizeMsg
	Program    *tea.Program
	Cursor     int
	Tr         *task.GormRepository
)