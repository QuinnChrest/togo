package main

import (
	"encoding/json"
	"log"
	"os"

	"togo/tui"
	constants "togo/tui/constants"
)

func getItemsFromFile() {
	var list []constants.Task

	// read in the contents of json file or create one if one doesn't exist
	file, err := os.Open("data.json")
	if err != nil {
		err := os.WriteFile("data.json", []byte(""), 0666)
		if err != nil {
			log.Fatal(err)
		} else {
			list = []constants.Task{}
		}
	} else {
		json.NewDecoder(file).Decode(&list)
	}
	defer file.Close()

	constants.List = list
}

func main() {
	getItemsFromFile()
	tui.Start()
}
