package entry

import {
	"fmt"
	"log"
	
	tea "github.com/charmbracelet/bubbletea"
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model.textInput, cmd = model.textInput.Update(msg)	
	return 	model, cmd
}

func (model Model) View() string {

}
