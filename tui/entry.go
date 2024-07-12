package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (model Model) EntryUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	model.TextInput, cmd = model.TextInput.Update(msg)
	return model, cmd
}

func (model Model) EntryView() string {
	return ""
}
