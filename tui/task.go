package tui

import (
	"strings"
	"togo/task"
	"togo/tui/constants"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	list []task.Task
	help help.Model
	page paginator.Model
}

// Initialize model for task view
func InitTask() tea.Model {
	// Get all tasks from sqlite db
	t, _ := constants.Tr.GetTasks()

	// Initialize paginator control
	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = constants.ActivePageDotStyle.Render("•")
	p.InactiveDot = constants.InactivePageDotStyle.Render("•")
	p.PerPage = (constants.WindowSize.Height - 8) / 3
	p.SetTotalPages(len(t))

	return Model{list: t, help: help.New(), page: p}
}

// Init run any intial IO on program start
func (m Model) Init() tea.Cmd {
	return nil
}

// Update loop ran every time an action occurs followed by a new render
func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		model.page.PerPage = (constants.WindowSize.Height - 7) / 3
		model.page.SetTotalPages(len(model.list))

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, taskKeyMap.Create):
			entry := InitEntry(task.Task{})
			return entry.Update(constants.WindowSize)

		case key.Matches(msg, taskKeyMap.Edit):
			entry := InitEntry(model.list[constants.Cursor+(model.page.Page*model.page.PerPage)])
			return entry.Update(constants.WindowSize)

		case key.Matches(msg, taskKeyMap.Quit):
			return model, tea.Quit

		case key.Matches(msg, taskKeyMap.Enter):
			constants.Tr.MarkComplete(&model.list[constants.Cursor+(model.page.Page*model.page.PerPage)])

		case key.Matches(msg, taskKeyMap.Delete):
			constants.Tr.DeleteTask(model.list[constants.Cursor+(model.page.Page*model.page.PerPage)].ID)
			model.list = removeTask(model.list, constants.Cursor+(model.page.Page*model.page.PerPage))
			if constants.Cursor != 0 {
				constants.Cursor--
			}

		case key.Matches(msg, taskKeyMap.Up):
			if constants.Cursor == 0 && model.page.Page != 0 {
				model.page.PrevPage()
				constants.Cursor = model.page.PerPage - 1
			} else if constants.Cursor != 0  {
				constants.Cursor--
			}

		case key.Matches(msg, taskKeyMap.Down):
			if constants.Cursor == model.page.PerPage-1 && model.page.Page != model.page.TotalPages-1 {
				model.page.NextPage()
				constants.Cursor = 0
			} else if constants.Cursor != model.page.PerPage-1 && constants.Cursor+(model.page.Page*model.page.PerPage) != len(model.list)-1 {
				constants.Cursor++
			}

		case key.Matches(msg, taskKeyMap.Left) && model.page.Page != 0:
			model.page.PrevPage()
			constants.Cursor = 0

		case key.Matches(msg, taskKeyMap.Right) && model.page.Page != model.page.TotalPages - 1:
			model.page.NextPage()
			constants.Cursor = 0
		}
	}

	return model, nil
}

// Render that occurs after every update loop
func (model Model) View() string {
	// Build string to be rendered for the user
	var b strings.Builder

	// Add header with some fancy lipgloss styling
	b.WriteString(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			constants.HeaderStyle.Render("——"),
			constants.HeaderTextStyle.Render("  ToGo  "),
			constants.HeaderStyle.Render(strings.Repeat("—", constants.WindowSize.Width)),
		) + "\n\n",
	)

	// Get pagination bounds for grabing the necessary slice
	start, end := model.page.GetSliceBounds(len(model.list))

	// Loop through slice and add list items with styling
	for i, v := range model.list[start:end] {
		var subtitle string

		if v.Complete {
			subtitle = "Completed - " + v.Time
		} else {
			subtitle = "Incomplete"
		}

		if i == constants.Cursor {
			b.WriteString(constants.SelectedItemStyle.Render(v.Description) + "\n")
			b.WriteString(constants.SelectedItemStyle.Render(subtitle) + "\n\n")
		} else {
			b.WriteString(constants.ItemStyle.Render(v.Description) + "\n")
			b.WriteString(constants.ItemStyle.Render(subtitle) + "\n\n")
		}
	}

	// If there were no list items added before add an info message stating there are no items to display
	if len(model.list) == 0 {
		b.WriteString(constants.GrayTextStyle.Render("There are no tasks to do yet. Add one by pressing 'c' to create.") + "\n\n")
	}

	// Add some new lines between the list items and footer menu so that the footer menu always sits on the bottom of the window
	b.WriteString(strings.Repeat("\n", getNewLines(len(model.list[start:end]))))

	// Add paginator control from Bubbles
	b.WriteString("  " + model.page.View() + "\n\n")

	// Add help bar explaining key controls
	// This is auto generated from Bubbles help control using our keymap
	b.WriteString(model.help.View(taskKeyMap))

	// Return the built view string
	return b.String()
}

// Get number of new lines needed to properly space the bottom menu to the bottom of the terminal
func getNewLines(listLen int) int {
	if listLen == 0 {
		return constants.WindowSize.Height-9
	} else {
		return constants.WindowSize.Height-(listLen*3+7)
	}
}

// Task has been removed from the DB but we need to remove the task from the local array
func removeTask(t []task.Task, index int) []task.Task {
	return append(t[:index], t[index+1:]...)
}

/* TASK VIEW KEY MAP ITEMS */

type taskkeymap struct {
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

var taskKeyMap = taskkeymap{
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
func (k taskkeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.Create, k.Edit, k.Delete, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k taskkeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{}, // first column
		{}, // second column
	}
}
