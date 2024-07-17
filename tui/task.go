package tui

import (
	"togo/task"
	"togo/tui/constants"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list []task.Task
}

func InitTask() tea.Model {
	// w, h, err := term.GetSize(int(os.Stdout.Fd()))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//vp := viewport.New(w, h)

	// pl, pc := getPages(h, len(items))

	t, _ := constants.Tr.GetTasks()

	return Model{
		list: t,
	}
}

// Init run any intial IO on program start
func (m Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Create):
			entry := InitEntry("")
			return entry.Update(constants.WindowSize)

		case key.Matches(msg, constants.Keymap.Edit):
			entry := InitEntry(model.list[constants.Cursor].Description)
			return entry.Update(constants.WindowSize)

		case key.Matches(msg, constants.Keymap.Up) && constants.Cursor != 0:
			constants.Cursor--

		case key.Matches(msg, constants.Keymap.Down) && constants.Cursor != len(model.list)-1:
			constants.Cursor++

		case key.Matches(msg, constants.Keymap.Quit):
			return model, tea.Quit
		}
	}

	return model, nil
}

func (model Model) View() string {
	content := ""

	for i, v := range model.list {
		if i == constants.Cursor {
			content += constants.SelectedItemStyle.Render(v.Description) + "\n"
		} else {
			content += constants.ItemStyle.Render(v.Description) + "\n"
		}
	}

	return content
}
