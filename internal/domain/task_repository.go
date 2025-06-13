package domain

type TaskRepository interface {
	Create(task *Task) error
	GetAll(filters Task) ([]Task, error)
	GetByID(id uint) (*Task, error)
	Update(task *Task) error
	Delete(id uint) error
}
