package service

import (
	"github.com/nintran52/my-tasks/model"
	"github.com/nintran52/my-tasks/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(r *repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) Create(task *model.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetByID(id string) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Update(task *model.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}
