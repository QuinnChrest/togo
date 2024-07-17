package tui

import (
	"log"
	"os"
	constants "togo/tui/constants"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type Entry struct {
	input	textinput.Model
}

func InitEntry(value string) *Entry {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	// initialize text input
	ti := textinput.New()
	ti.Placeholder = "Add Task"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = w

	if len(value) > 0 {
		ti.SetValue(value)
	}

	m := Entry{input: ti}

	return &m
}

func (model Entry) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (model Entry) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Back):
			return InitTask(), nil
		case key.Matches(msg, constants.Keymap.Quit):
			return model, tea.Quit
		}
	}

	model.input, cmd = model.input.Update(msg)
	return model, cmd
}

func (model Entry) View() string {
	return model.input.View()
}
