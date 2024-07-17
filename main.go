package main

import (
	"fmt"
	"log"

	task "togo/task"
	"togo/tui"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openSqlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("new.db"), &gorm.Config{})
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
	db, err := openSqlite()
	if err != nil {
		log.Fatal(err)
	}
	tr := task.GormRepository{DB: db}
	tui.Start(tr)
}
