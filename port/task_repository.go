package port

import "github.com/nintran52/my-tasks/domain"

type TaskRepository interface {
	Create(task *domain.Task) error
	FindAll() ([]domain.Task, error)
	FindByID(id uint) (*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id uint) error
}
