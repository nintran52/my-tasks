package domain

type TaskRepository interface {
	Create(task *Task) error
	GetAll(limit, offset int, filters Task) ([]Task, int64, error)
	GetByID(id uint) (*Task, error)
	Update(task *Task) error
	Delete(id uint) error
}
