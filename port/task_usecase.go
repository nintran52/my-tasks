package port

import "github.com/nintran52/my-tasks/domain"

type TaskService interface {
	CreateTask(task *domain.Task) error
	GetTasks() ([]domain.Task, error)
	GetTask(id uint) (*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id uint) error
}
