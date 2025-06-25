package domain_test

import (
	"testing"

	"github.com/nintran52/my-tasks/internal/domain"
)

func TestValidateTask(t *testing.T) {
	tests := []struct {
		name    string
		task    domain.Task
		wantErr bool
	}{
		{
			name: "valid task",
			task: domain.Task{
				Title: "Valid Task",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			task: domain.Task{
				Title: "",
			},
			wantErr: true,
		},
		{
			name: "short title",
			task: domain.Task{
				Title: "No",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.task.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
