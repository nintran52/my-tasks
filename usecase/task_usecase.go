package usecase

import (
	"github.com/nintran52/my-tasks/domain"
	"github.com/nintran52/my-tasks/port"
)

type taskService struct {
	repo port.TaskRepository
}

func NewTaskService(repo port.TaskRepository) port.TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *domain.Task) error {
	return s.repo.Create(task)
}
func (s *taskService) GetTasks() ([]domain.Task, error) {
	return s.repo.FindAll()
}
func (s *taskService) GetTask(id uint) (*domain.Task, error) {
	return s.repo.FindByID(id)
}
func (s *taskService) UpdateTask(task *domain.Task) error {
	return s.repo.Update(task)
}
func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
