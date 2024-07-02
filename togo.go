package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	header = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("63"))

	footer = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("63"))

	item = lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("7")).
		BorderLeft(true).
		PaddingLeft(2)

	selected = item.
		Foreground(lipgloss.Color("63")).
		BorderForeground(lipgloss.Color("63")).
		Bold(true)
)

type Model struct {
	Items     []Task
	cursor    int
	add       bool
	textInput textinput.Model
	w, h      int
}

type Task struct {
	Description string
	Checked     bool
}

func initialModel() Model {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	// initialize text input
	ti := textinput.New()
	ti.Placeholder = "Add Task"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = w

	return Model{
		Items:     getItemsFromFile(),
		add:       false,
		textInput: ti,
		w:         w,
		h:         h,
	}
}

func getItemsFromFile() []Task {
	var list []Task

	// read in the contents of json file or create one if one doesn't exist
	file, err := os.Open("data.json")
	if err != nil {
		err := os.WriteFile("data.json", []byte(""), 0666)
		if err != nil {
			log.Fatal(err)
		} else {
			list = []Task{}
		}
	} else {
		json.NewDecoder(file).Decode(&list)
	}
	defer file.Close()

	return list
}

func addItem(model *Model) {
	model.add = false
	model.Items = append(model.Items, Task{Description: model.textInput.Value(), Checked: false})
	model.textInput.SetValue("")
}

func (model Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "esc":
			if !model.add {
				saveTasks(model)
				return model, tea.Quit
			} else {
				model.textInput.SetValue("")
				model.add = false
			}

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if model.cursor > 0 && !model.add {
				model.cursor--
			}

		case "enter":
			if !model.add {
				model.Items[model.cursor].Checked = !model.Items[model.cursor].Checked
			} else {
				addItem(&model)
			}

		case "a":
			if !model.add {
				model.add = true
				model.textInput, cmd = model.textInput.Update(tea.KeyMsg{})
				return model, cmd
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if model.cursor < len(model.Items)-1 && !model.add {
				model.cursor++
			}
		}

	case tea.WindowSizeMsg:
		if model.h != msg.Height || model.w != msg.Width {
			model.h, model.w = msg.Height, msg.Width

			model.textInput.Width = model.w
		}
	}

	if !model.add {
		return model, nil
	} else {
		model.textInput, cmd = model.textInput.Update(msg)
		return model, cmd
	}
}

func (model Model) View() string {
	output := header.Width(model.w).Render("ToGo") + "\n\n"

	for i, v := range model.Items{
		var content string

		if i == model.cursor {
			content += selected.Width(model.w).Render(v.Description) + "\n"
			content += selected.Width(model.w).Render(strconv.FormatBool(v.Checked)) + "\n"
		} else {
			content += item.Width(model.w).Render(v.Description) + "\n"
			content += item.Width(model.w).Render(strconv.FormatBool(v.Checked)) + "\n"
		}

		output += content + "\n"
	}

	// The footer
	output += footer.Width(model.w).Render("Press Esc to quit.")

	// Send the UI for rendering
	return output
}

func saveTasks(model Model) {
	jsonData, err := json.MarshalIndent(model.Items, "", "      ")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("data.json", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
