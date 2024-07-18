package task

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	Complete    bool
	Time        string
}

type Repository interface {
	DeleteTask(taskID uint) error
	GetTasks() ([]Task, error)
	CreateTask(description []byte) error
	EditTask(id uint, description string) error
}

type GormRepository struct {
	DB *gorm.DB
}

func (g *GormRepository) DeleteTask(taskID uint) error {
	result := g.DB.Delete(&Task{}, taskID)
	return result.Error
}

func (g *GormRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	if err := g.DB.Find(&tasks).Error; err != nil {
		return tasks, fmt.Errorf("table is empty: %v", err)
	}
	return tasks, nil
}

func (g *GormRepository) CreateTask(description string) error {
	task := Task{Description: description}
	result := g.DB.Create(&task)
	return result.Error
}

func (g *GormRepository) EditTask(t Task) error {
	if err := g.DB.Save(&t).Error; err != nil {
		return fmt.Errorf("unable to save task: %w", err)
	}
	return nil
}

func (g *GormRepository) MarkComplete(t *Task) error {
	t.Complete = !t.Complete
	t.Time = time.Now().Format("01/02/2006 03:04 PM")
	if err := g.DB.Save(&t).Error; err != nil {
		return fmt.Errorf("unable to save task: %w", err)
	}
	return nil
}
