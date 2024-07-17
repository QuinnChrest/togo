package constants

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("63")).
			PaddingLeft(2).
			PaddingRight(2)

	FooterStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			MarginLeft(1)

	ItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("7")).
			BorderLeft(true).
			PaddingLeft(2)

	SelectedItemStyle = ItemStyle.
				Foreground(lipgloss.Color("63")).
				BorderForeground(lipgloss.Color("63")).
				Bold(true)
)

var (
	WindowSize 	tea.WindowSizeMsg
	Program		*tea.Program
	List		[]Task
)

type keymap struct {
	Create	key.Binding
	Edit	key.Binding
	Delete	key.Binding
	Quit	key.Binding
	Back	key.Binding
	Up		key.Binding
	Down	key.Binding
}

var Keymap = keymap{
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "create"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "create"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Up: key.NewBinding(
		key.WithKeys("j"),
		key.WithHelp("j", "Up"),
	),
	Down: key.NewBinding(
		key.WithKeys("k"),
		key.WithHelp("k", "Down"),
	),
}

type Task struct {
	Description string
	Complete    bool
	Time        string
}
