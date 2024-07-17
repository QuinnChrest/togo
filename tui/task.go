package tui

import (
	constants "togo/tui/constants"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	nav mode = iota
	edit
	create
)

type Model struct {
	mode		mode
	list		[]constants.Task
	cursor		int
}

func InitTask() tea.Model {
	// w, h, err := term.GetSize(int(os.Stdout.Fd()))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//vp := viewport.New(w, h)

	// pl, pc := getPages(h, len(items))
 
	return Model{
		list:		constants.List,
		mode: 		nav,
		cursor:		0,
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
			entry := InitEntry(model.list[model.cursor].Description)
			return entry.Update(constants.WindowSize)

		case key.Matches(msg, constants.Keymap.Up) && model.cursor != 0:
			model.cursor--

		case key.Matches(msg, constants.Keymap.Down) && model.cursor != len(model.list) - 1:
			model.cursor++

		case key.Matches(msg, constants.Keymap.Quit):
			return model, tea.Quit
		}
	}

	return model, nil
}

func (model Model) View() string {
	content := ""

	for i, v := range model.list {
		if i == model.cursor {
			content += constants.SelectedItemStyle.Render(v.Description) + "\n"
		} else {
			content += constants.ItemStyle.Render(v.Description) + "\n"
		}
	}

	return content
}
