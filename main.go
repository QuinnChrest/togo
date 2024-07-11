package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	. "togo/tui/constants"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

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

	vp := viewport.New(w, h)

	items := getItemsFromFile()

	pl, pc := getPages(h, len(items))

	return Model{
		Items:      items,
		State:      ViewState,
		TextInput:  ti,
		ViewPort:   vp,
		Width:      w,
		Height:     h,
		Page:       0,
		PageLength: pl,
		PageCount:  pc,
	}
}

func getPages(height int, itemCount int) (pl int, pc int) {
	if height < 7 || itemCount == 0 {
		return 0, 0
	}

	itemsPerPage := int(math.Floor(float64(height-4) / float64(3)))

	return itemsPerPage, int(math.Ceil(float64(itemCount) / float64(itemsPerPage)))
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
	model.Items = append(model.Items, Task{Description: model.TextInput.Value(), Complete: false, Time: time.Now().Format("01/02/2006 03:04 PM")})
	model.TextInput.SetValue("")
}

func removeItem(slice []Task, s int) []Task {
	return append(slice[:s], slice[s+1:]...)
}

func (model Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch model.State {
		case ViewState:
			return model, nil
		}

	case tea.WindowSizeMsg:
		if model.h != msg.Height || model.w != msg.Width {
			model.h, model.w = msg.Height, msg.Width

			model.viewPort.Width, model.viewPort.Height = model.w, model.h

			model.textInput.Width = model.w

			model.pl, model.pc = getPages(model.h, len(model.Items))

			if model.pc > model.p {
				model.p = model.pc - 1
			}
		}
	}

	return model, nil
}

func (model Model) View() string {
	output := HeaderStyle.Render("ToGo - "+strconv.Itoa(len(model.Items))+" items") + "\n\n"

	if !false {
		page := model.Items[model.p*model.pl : min(((model.p*model.pl)+model.pl), len(model.Items))]

		for i, v := range page {
			var content, subtitle string

			if v.Complete {
				subtitle += "Completed - " + v.Time
			} else {
				subtitle += "Incomplete"
			}

			if i == model.Cursor {
				content += SelectedItemStyle.Width(model.Width).Render(v.Description) + "\n"
				content += SelectedItemStyle.Width(model.Width).Render(subtitle) + "\n"
			} else {
				content += SelectedItemStyle.Width(model.Width).Render(v.Description) + "\n"
				content += SelectedItemStyle.Width(model.Width).Render(subtitle) + "\n"
			}

			output += content + "\n"
		}

		output += FooterStyle.MarginTop(model.Height - ((len(page) * 3) + 2) - 2).Render(getFooter(model))
	} else {
		output += model.TextInput.View()
	}

	model.ViewPort.SetContent(output)

	// Send the UI for rendering
	return model.ViewPort.View()
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getFooter(model Model) string {
	return getPaginator(model.Page, model.PageCount) + getControls()
}

func getPaginator(page, pageCount int) string {
	content := ""

	if pageCount == 1 {
		return content + "\n"
	}

	for i := range pageCount {
		if i == page {
			content += "•"
		} else {
			content += "•"
		}
	}

	return content + "\n"
}

func getControls() string {
	return fmt.Sprintf(
		"%s %s %s %s %s %s %s %s",
		"enter",
		"mark complete •",
		"a",
		"add •",
		"delete",
		"remove •",
		"q/escape",
		"quit",
	)
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
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
