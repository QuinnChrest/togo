package constants

import {
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
}

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

	SelectedItemStyle = item.
			Foreground(lipgloss.Color("63")).
			BorderForeground(lipgloss.Color("63")).
			Bold(true)
)

var (
	ViewState := 1
	CreateState := 2
	EditState := 3
)

type keymap struct {
	Create	key.Binding
	Edit	key.Binding
	Delete	key.Binding
	Quit	key.Binding
	Back	key.Binding
}

var Keymap = keymap{
	Create: key.NewBinding(
		key.WithKeys("c")
		key.WithHelp("c", "create")
	),
	Edit: key.NewBinding(
		key.WithKeys("e")
		key.WithHelp("e", "edit")
	),
	Delete: key.NewBinding(
		key.WithKeys("d")
		key.WithHelp("d", "delete")
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q")
		key.WithHelp("ctrl+c/q", "create")
	),
	Back: key.NewBinding(
		key.WithKeys("esc")
		key.WithHelp("esc", "back")
	)
}