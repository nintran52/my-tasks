package repository

import (
	"github.com/nintran52/my-tasks/internal/domain/model"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error { return r.db.Create(task).Error }
func (r *taskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	return tasks, r.db.Find(&tasks).Error
}
func (r *taskRepository) FindByID(id uint) (*model.Task, error) {
	var task model.Task
	result := r.db.First(&task, id)
	return &task, result.Error
}
func (r *taskRepository) Update(task *model.Task) error { return r.db.Save(task).Error }
func (r *taskRepository) Delete(id uint) error          { return r.db.Delete(&model.Task{}, id).Error }
