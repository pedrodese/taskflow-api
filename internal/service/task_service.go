package service

import (
	"taskFlow/internal/repository"
	"taskFlow/internal/schema"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetAllTasks() ([]schema.Task, error) {
	return s.Repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id string) (*schema.Task, error) {
	return s.Repo.GetTaskByID(id)
}

func (s *TaskService) CreateTask(task *schema.Task) error {
	return s.Repo.CreateTask(task)
}

func (s *TaskService) UpdateTask(task *schema.Task) error {
	return s.Repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id string) error {
	return s.Repo.DeleteTask(id)
}
