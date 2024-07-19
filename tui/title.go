package tui

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"togo/tui/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	url = lipgloss.NewStyle().Foreground(special).Render

	descStyle = lipgloss.NewStyle().MarginTop(1)

	titleStyle = lipgloss.NewStyle().
		MarginLeft(1).
		MarginRight(5).
		Padding(0, 4).
		Italic(true).
		Bold(true).
		Foreground(lipgloss.Color("#FFF7DB")).
		SetString("To Go")

	infoStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(subtle)

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()
)

type Title struct{
	timer *time.Timer
	exit bool
}

func InitTitle() tea.Model {
	// Get terminal window size right off the bat for certain spacing needs
	var err error
	constants.WindowSize.Width, constants.WindowSize.Height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	
	return Title{}
}

// Init run any intial IO on program start
func (t Title) Init() tea.Cmd {
	return nil
}

// Update loop ran every time an action occurs followed by a new render
func (t Title) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	t.timer = time.NewTimer(time.Second * 2)

	

	return t, nil
}

func (t Title) View() string {
	var (
		colors []lipgloss.TerminalColor
		title  strings.Builder
	)

	colors = append(
		colors,
		lipgloss.Color("#ff42d0"),
		lipgloss.Color("#ff42ff"),
		lipgloss.Color("#ec42ff"),
		lipgloss.Color("#d042ff"),
		lipgloss.Color("#a142ff"),
	)

	for i, v := range colors {
		const offset = 2
		fmt.Fprint(&title, titleStyle.MarginLeft(i*offset).Background(v))
		if i < len(colors)-1 {
			title.WriteRune('\n')
		}
	}

	desc := lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("A Go based To Do list application (cause why not)"),
		infoStyle.Render("By Quinn Chrest"+divider+url("https://github.com/QuinnChrest/togo")),
	)

	row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
	
	return lipgloss.NewStyle().MarginTop((constants.WindowSize.Height - 5) / 2).MarginLeft((constants.WindowSize.Width - lipgloss.Width(row)) / 2).Render(row)
}
