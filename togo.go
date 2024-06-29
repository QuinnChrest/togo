package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
    header = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FAFAFA")).
        Background(lipgloss.Color("#7D56F4")).
        Height(1)
)

type Model struct {
    Items []Task
    cursor int
    add bool
    textInput textinput.Model
    viewPort viewport.Model
}

type Task struct {
    Description string
    Checked bool
}

func initialModel() Model{
    // initialize text input
    ti := textinput.New()
	ti.Placeholder = "Add Task"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

    w, h, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        log.Fatal(err)
    }

    vp := viewport.New(w, h-1)
    vp.YPosition = 2

    return Model{
        Items: getItemsFromFile(),
        add: false,
        textInput: ti,
        viewPort: vp,
    }
}

func getItemsFromFile() []Task{
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
            if model.viewPort.Height != msg.Height - 1 || model.viewPort.Width != msg.Width {
                model.viewPort.Height = msg.Height - 1
                model.viewPort.Width = msg.Width

                model.textInput.Width = model.viewPort.Width
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
    output := header.Width(model.viewPort.Width).Render("ToGo")

    if !model.add {
        content := ""

        // Iterate over our choices
        for i, item := range model.Items {

            // Is the cursor pointing at this choice?
            cursor := " " // no cursor
            if model.cursor == i {
                cursor = ">" // cursor!
            }
    
            checked := " "
            if(item.Checked){
                checked = "x"
            }
    
            // Render the row
            content += fmt.Sprintf("%s [%s] %s", cursor, checked, item.Description)
            if i < len(model.Items) - 1 {
                content += "\n"
            }
        }

        model.viewPort.SetContent(content)
        output += model.viewPort.View()
    } else {
        output += model.textInput.View()
    }

    // The footer
    output += header.Width(model.viewPort.Width).Render("Press Esc to quit.")

    // Send the UI for rendering
    return output
}

func saveTasks(model Model){
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