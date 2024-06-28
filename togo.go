package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

type Model struct {
    Items []Task
    Cursor int
    ViewPort viewport.Model
}

type Task struct {
    Description string
    Checked bool
}

func (model Model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            saveTasks(model)
            return model, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if model.Cursor > 0 {
                model.Cursor--
            }

        case "enter":
            model.Items[model.Cursor].Checked = !model.Items[model.Cursor].Checked

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if model.Cursor < len(model.Items)-1 {
                model.Cursor++
            }
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return model, nil
}

func (model Model) View() string {
    // The header
    s := "What should we buy at the market?\n\n"

    // Iterate over our choices
    for i, item := range model.Items {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if model.Cursor == i {
            cursor = ">" // cursor!
        }

        checked := " "
        if(item.Checked){
            checked = "x"
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item.Description)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func (model Model) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, model.ViewPort.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (model Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", model.ViewPort.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, model.ViewPort.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
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
    m := Model{}

    // read in the contents of json file or create one if one doesn't exist
    file, err := os.Open("data.json")
    if err != nil {
        err := os.WriteFile("data.json", []byte(""), 0666)
        if err != nil {
            log.Fatal(err)
        } else {
            m.Items = []Task{}
        }
    } else {
        json.NewDecoder(file).Decode(&m.Items)
    }
    file.Close()

    p := tea.NewProgram(m)
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}