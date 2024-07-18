package constants

import (
	"togo/task"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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

var (
	WindowSize tea.WindowSizeMsg
	Program    *tea.Program
	Cursor     int
	Tr         *task.GormRepository
)

type keymap struct {
	Create key.Binding
	Edit   key.Binding
	Delete key.Binding
	Quit   key.Binding
	Enter  key.Binding
	Up     key.Binding
	Down   key.Binding
	Right  key.Binding
	Left   key.Binding
}

var Keymap = keymap{
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Create"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "Edit"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "Delete"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "Quit"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Mark Complete"),
	),
	Up: key.NewBinding(
		key.WithKeys("j", "up"),
		key.WithHelp("j/↑", "Up"),
	),
	Down: key.NewBinding(
		key.WithKeys("k", "down"),
		key.WithHelp("k/↓", "Down"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("", ""),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("", ""),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.Create, k.Edit, k.Delete, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{}, // first column
		{}, // second column
	}
}
