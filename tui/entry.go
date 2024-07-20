package tui

import (
	"togo/task"
	"togo/tui/constants"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Entry struct {
	input textinput.Model
	task  task.Task
	help  help.Model
}

// Initialize entry struct for text input view
func InitEntry(task task.Task) *Entry {
	// initialize text input
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = int(float64(constants.WindowSize.Width) * 0.6)
	ti.SetValue(task.Description)

	return &Entry{task: task, help: help.New(), input: ti}
}

func (model Entry) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// Update loop ran every time an action occurs followed by a new render
func (model Entry) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, entryKeyMap.Back):
			return InitTask(), nil
		case key.Matches(msg, entryKeyMap.Quit):
			return model, tea.Quit
		case key.Matches(msg, entryKeyMap.Enter):
			if model.task.ID != 0 {
				model.task.Description = model.input.Value()
				constants.Tr.EditTask(model.task)
			} else {
				constants.Tr.CreateTask(model.input.Value())
			}
			return InitTask(), nil
		}
	}

	model.input, cmd = model.input.Update(msg)
	return model, cmd
}

// Render that occurs after every update loop
func (model Entry) View() string {
	return lipgloss.NewStyle().
		Width(constants.WindowSize.Width).
		Height(constants.WindowSize.Height).
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				"  Add New Task",
				lipgloss.NewStyle().
					Border(lipgloss.RoundedBorder()).
					BorderForeground(lipgloss.Color("#ec42ff")).
					Render(" "+model.input.View()),
				model.help.View(entryKeyMap),
			),
		)
}

/* ENTRY VIEW KEY MAP ITEMS */

type entrykeymap struct {
	Quit  key.Binding
	Back  key.Binding
	Enter key.Binding
}

var entryKeyMap = entrykeymap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "Quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Back"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Create"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k entrykeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Back, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k entrykeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{}, // first column
		{}, // second column
	}
}
