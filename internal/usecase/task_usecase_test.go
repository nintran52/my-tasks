package usecase_test

import (
	"testing"

	"github.com/nintran52/my-tasks/internal/domain"
	"github.com/nintran52/my-tasks/internal/domain/mocks"
	"github.com/nintran52/my-tasks/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockRepo := &mocks.TaskRepository{}
	uc := usecase.NewTaskUsecase(mockRepo)

	cases := []struct {
		name    string
		task    *domain.Task
		wantErr bool
	}{
		{
			name:    "Valid Task",
			task:    &domain.Task{ID: 1, Title: "Valid Task", Done: false},
			wantErr: false,
		},
		{
			name:    "Invalid Task",
			task:    &domain.Task{ID: 2, Title: "", Done: false}, // Assuming Title is required
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				mockRepo.On("Create", c.task).Return(assert.AnError)
			} else {
				mockRepo.On("Create", c.task).Return(nil)
			}
			err := uc.CreateTask(c.task)
			if c.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTasks(t *testing.T) {
	mockRepo := &mocks.TaskRepository{}
	uc := usecase.NewTaskUsecase(mockRepo)

	cases := []struct {
		name    string
		limit   int
		offset  int
		filters domain.Task
		wantErr bool
		wantLen int
	}{
		{
			name:    "Valid Request",
			limit:   10,
			offset:  0,
			filters: domain.Task{Done: false},
			wantErr: false,
			wantLen: 2,
		},
		{
			name:    "Invalid Request",
			limit:   -1, // Invalid limit
			offset:  0,
			filters: domain.Task{},
			wantErr: true,
			wantLen: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				mockRepo.On("GetAll", c.limit, c.offset, c.filters).Return(nil, int64(0), assert.AnError)
			} else {
				mockRepo.On("GetAll", c.limit, c.offset, c.filters).Return([]domain.Task{{ID: 1}, {ID: 2}}, int64(2), nil)
			}
			tasks, count, err := uc.GetTasks(c.limit, c.offset, c.filters)
			if c.wantErr {
				assert.Error(t, err)
				assert.Equal(t, 0, len(tasks))
				assert.Equal(t, int64(0), count)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.wantLen, len(tasks))
				assert.Equal(t, int64(2), count)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	mockRepo := &mocks.TaskRepository{}
	uc := usecase.NewTaskUsecase(mockRepo)

	cases := []struct {
		name    string
		task    *domain.Task
		wantErr bool
	}{
		{
			name:    "Valid Task Update",
			task:    &domain.Task{ID: 1, Title: "Updated Task", Done: true},
			wantErr: false,
		},
		{
			name:    "Invalid Task Update",
			task:    &domain.Task{ID: 2, Title: "", Done: false}, // Assuming Title is required
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				mockRepo.On("Update", c.task).Return(assert.AnError)
			} else {
				mockRepo.On("Update", c.task).Return(nil)
			}
			err := uc.UpdateTask(c.task)
			if c.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	mockRepo := &mocks.TaskRepository{}
	uc := usecase.NewTaskUsecase(mockRepo)

	cases := []struct {
		name    string
		id      uint
		wantErr bool
	}{
		{
			name:    "Valid ID",
			id:      1,
			wantErr: false,
		},
		{
			name:    "Invalid ID",
			id:      999, // Assuming this ID does not exist
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				mockRepo.On("Delete", c.id).Return(assert.AnError)
			} else {
				mockRepo.On("Delete", c.id).Return(nil)
			}
			err := uc.DeleteTask(c.id)
			if c.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
