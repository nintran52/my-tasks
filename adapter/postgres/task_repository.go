package postgres

import (
	"github.com/nintran52/my-tasks/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}
func (r *taskRepository) FindAll() ([]domain.Task, error) {
	var tasks []domain.Task
	return tasks, r.db.Find(&tasks).Error
}
func (r *taskRepository) FindByID(id uint) (*domain.Task, error) {
	var task domain.Task
	result := r.db.First(&task, id)
	return &task, result.Error
}
func (r *taskRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}
func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Task{}, id).Error
}
