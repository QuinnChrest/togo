package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (model Model) TaskUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model Model) TaskView() string {
	return ""
}
