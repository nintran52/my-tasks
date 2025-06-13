package usecase

import "github.com/nintran52/my-tasks/internal/domain"

type TaskUsecase struct {
	Repo domain.TaskRepository
}

func NewTaskUsecase(r domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{Repo: r}
}

func (uc *TaskUsecase) CreateTask(t *domain.Task) error {
	return uc.Repo.Create(t)
}

func (uc *TaskUsecase) GetTasks() ([]domain.Task, error) {
	return uc.Repo.GetAll()
}

func (uc *TaskUsecase) GetTaskByID(id uint) (*domain.Task, error) {
	return uc.Repo.GetByID(id)
}

func (uc *TaskUsecase) UpdateTask(t *domain.Task) error {
	return uc.Repo.Update(t)
}

func (uc *TaskUsecase) DeleteTask(id uint) error {
	return uc.Repo.Delete(id)
}
