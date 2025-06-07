package service

import (
	"github.com/nintran52/my-tasks/internal/domain/model"
	"github.com/nintran52/my-tasks/internal/domain/port"
)

type taskService struct {
	repo port.TaskRepository
}

func NewTaskService(repo port.TaskRepository) port.TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *model.Task) error    { return s.repo.Create(task) }
func (s *taskService) GetTasks() ([]model.Task, error)      { return s.repo.FindAll() }
func (s *taskService) GetTask(id uint) (*model.Task, error) { return s.repo.FindByID(id) }
func (s *taskService) UpdateTask(task *model.Task) error    { return s.repo.Update(task) }
func (s *taskService) DeleteTask(id uint) error             { return s.repo.Delete(id) }
