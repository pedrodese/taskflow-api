package repository

import (
	"taskFlow/internal/schema"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetAllTasks() ([]schema.Task, error) {
	var tasks []schema.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(id string) (*schema.Task, error) {
	var task schema.Task
	if err := r.DB.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) CreateTask(task *schema.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) UpdateTask(task *schema.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) DeleteTask(id string) error {
	return r.DB.Delete(&schema.Task{}, "id = ?", id).Error
}
