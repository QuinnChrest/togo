package task

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// DB gorm model
type Task struct {
	gorm.Model
	Description string
	Complete    bool
	Time        string
}

// Data repository interface
type Repository interface {
	DeleteTask(taskID uint) error
	GetTasks() ([]Task, error)
	CreateTask(description []byte) error
	EditTask(id uint, description string) error
}

// Data struct containing memory reference to database
type GormRepository struct {
	DB *gorm.DB
}

// Delete task by ID
func (g *GormRepository) DeleteTask(taskID uint) error {
	result := g.DB.Delete(&Task{}, taskID)
	return result.Error
}

// Get all tasks
func (g *GormRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	if err := g.DB.Find(&tasks).Error; err != nil {
		return tasks, fmt.Errorf("table is empty: %v", err)
	}
	return tasks, nil
}

// Create a new task
func (g *GormRepository) CreateTask(description string) error {
	task := Task{Description: description}
	result := g.DB.Create(&task)
	return result.Error
}

// Edit an existing task
func (g *GormRepository) EditTask(t Task) error {
	if err := g.DB.Save(&t).Error; err != nil {
		return fmt.Errorf("unable to save task: %w", err)
	}
	return nil
}

// Mark a task as complete
func (g *GormRepository) MarkComplete(t *Task) error {
	t.Complete = !t.Complete
	t.Time = time.Now().Format("01/02/2006 03:04 PM")
	if err := g.DB.Save(&t).Error; err != nil {
		return fmt.Errorf("unable to save task: %w", err)
	}
	return nil
}
