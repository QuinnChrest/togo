package task

import (
	"fmt"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	Complete    bool
	Time        string
}

// Repository the CRUD functionality for entries
type Repository interface {
	DeleteTask(taskID uint) error
	GetTasks() ([]Task, error)
	CreateTask(description []byte) error
	EditTask(id uint, description string) error
}

// GormRepository holds the gorm DB and is a TaskRepository
type GormRepository struct {
	DB *gorm.DB
}

// DeleteEntryByID delete an entry by its ID
func (g *GormRepository) DeleteTask(taskID uint) error {
	result := g.DB.Delete(&Task{}, taskID)
	return result.Error
}

// GetEntriesByProjectID get all entries for a given project
func (g *GormRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	if err := g.DB.Find(&tasks).Error; err != nil {
		return tasks, fmt.Errorf("Table is empty: %v", err)
	}
	return tasks, nil
}

// CreateEntry create a new entry in the database
func (g *GormRepository) CreateTask(description []byte) error {
	task := Task{Description: string(description[:])}
	result := g.DB.Create(&task)
	return result.Error
}

// RenameProject rename an existing project
func (g *GormRepository) EditTask(id uint, description string) error {
	var task Task
	if err := g.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return fmt.Errorf("unable to edit task: %w", err)
	}
	task.Description = description
	if err := g.DB.Save(&task).Error; err != nil {
		return fmt.Errorf("unable to save task: %w", err)
	}
	return nil
}
