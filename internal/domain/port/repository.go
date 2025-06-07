package port

import "github.com/nintran52/my-tasks/internal/domain/model"

type TaskRepository interface {
	Create(task *model.Task) error
	FindAll() ([]model.Task, error)
	FindByID(id uint) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id uint) error
}
