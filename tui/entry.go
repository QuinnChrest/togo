package tui

import (
	. "togo/tui/constants"

	tea "github.com/charmbracelet/bubbletea"
)

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	model.textInput, cmd = model.textInput.Update(msg)
	return model, cmd
}

func (model Model) View() string {

}
