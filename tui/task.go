package tui

import (
	"log"
	"os"
	"strings"
	"togo/task"
	"togo/tui/constants"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type Model struct {
	list []task.Task
	help help.Model
	page paginator.Model
}

func InitTask() tea.Model {
	var err error
	constants.WindowSize.Width, constants.WindowSize.Height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	t, _ := constants.Tr.GetTasks()

	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = constants.ActivePageDotStyle.Render("•")
	p.InactiveDot = constants.InactivePageDotStyle.Render("•")
	p.PerPage = (constants.WindowSize.Height - 8) / 3
	p.SetTotalPages(len(t))

	return Model{
		list: t,
		help: help.New(),
		page: p,
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
		model.page.PerPage = (constants.WindowSize.Height - 7) / 3
		model.page.SetTotalPages(len(model.list))

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Create):
			entry := InitEntry(task.Task{})
			return entry.Update(constants.WindowSize)

		case key.Matches(msg, constants.Keymap.Edit):
			entry := InitEntry(model.list[constants.Cursor+(model.page.Page*model.page.PerPage)])
			return entry.Update(constants.WindowSize)

		case key.Matches(msg, constants.Keymap.Quit):
			return model, tea.Quit

		case key.Matches(msg, constants.Keymap.Enter):
			constants.Tr.MarkComplete(&model.list[constants.Cursor+(model.page.Page*model.page.PerPage)])

		case key.Matches(msg, constants.Keymap.Delete):
			constants.Tr.DeleteTask(model.list[constants.Cursor+(model.page.Page*model.page.PerPage)].ID)
			model.list = removeTask(model.list, constants.Cursor+(model.page.Page*model.page.PerPage))
			if constants.Cursor != 0 {
				constants.Cursor--
			}

		case key.Matches(msg, constants.Keymap.Up) && constants.Cursor != 0:
			constants.Cursor--

		case key.Matches(msg, constants.Keymap.Down) && constants.Cursor != len(model.list)-1:
			constants.Cursor++

		case key.Matches(msg, constants.Keymap.Left):
			model.page.PrevPage()

		case key.Matches(msg, constants.Keymap.Right):
			model.page.NextPage()
		}
	}

	return model, nil
}

func (model Model) View() string {
	var b strings.Builder

	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, constants.HeaderStyle.Render("——"), constants.HeaderTextStyle.Render("  ToGo  "), constants.HeaderStyle.Render(strings.Repeat("—", constants.WindowSize.Width))) + "\n\n")

	start, end := model.page.GetSliceBounds(len(model.list))

	for i, v := range model.list[start:end] {
		var subtitle string

		if v.Complete {
			subtitle += "Completed - " + v.Time
		} else {
			subtitle += "Incomplete"
		}

		if i == constants.Cursor {
			b.WriteString(constants.SelectedItemStyle.Render(v.Description) + "\n")
			b.WriteString(constants.SelectedItemStyle.Render(subtitle) + "\n\n")
		} else {
			b.WriteString(constants.ItemStyle.Render(v.Description) + "\n")
			b.WriteString(constants.ItemStyle.Render(subtitle) + "\n\n")
		}
	}

	if len(model.list) == 0 {
		b.WriteString("There are no items to do yet. Add one by pressing 'c' to create.\n\n")
	}

	b.WriteString(strings.Repeat("\n", constants.WindowSize.Height-(len(model.list[start:end])*3+7)))

	b.WriteString("  " + model.page.View() + "\n\n")

	b.WriteString(model.help.View(constants.Keymap))

	return b.String()
}

func removeTask(t []task.Task, index int) []task.Task {
	return append(t[:index], t[index+1:]...)
}
