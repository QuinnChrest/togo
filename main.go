package main

import (
	"fmt"
	"log"

	"togo/task"
	"togo/tui"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Open our sqlite DB with the gorm package
func openSqlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("togo.db"), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("unable to open database: %w", err)
	}
	err = db.AutoMigrate(&task.Task{})
	if err != nil {
		return db, fmt.Errorf("unable to migrate database: %w", err)
	}
	return db, nil
}

func main() {
	// Grab our gorm DB and start the Bubble Tea Terminal UI
	db, err := openSqlite()
	if err != nil {
		log.Fatal(err)
	}
	tr := task.GormRepository{DB: db}
	tui.Start(tr)
}
