package tui

import (
	"togo/task"
	"togo/tui/constants"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Entry struct {
	input textinput.Model
	edit  bool
	task  task.Task
}

func InitEntry(task task.Task) *Entry {
	// initialize text input
	ti := textinput.New()
	ti.Placeholder = "Add Task"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = constants.WindowSize.Width

	m := Entry{task: task}

	if task.Description != "" {
		m.edit = true
		ti.SetValue(task.Description)
	}

	m.input = ti

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
		case key.Matches(msg, constants.Keymap.Enter):
			if model.edit {
				constants.Tr.EditTask(model.task.ID, model.input.Value())
			} else {
				constants.Tr.CreateTask([]byte(model.input.Value()))
			}
			return InitTask(), nil
		}
	}

	model.input, cmd = model.input.Update(msg)
	return model, cmd
}

func (model Entry) View() string {
	return model.input.View()
}
