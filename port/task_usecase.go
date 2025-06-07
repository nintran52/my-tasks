package port

import "github.com/nintran52/my-tasks/internal/domain/model"

type TaskService interface {
	CreateTask(task *model.Task) error
	GetTasks() ([]model.Task, error)
	GetTask(id uint) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id uint) error
}
