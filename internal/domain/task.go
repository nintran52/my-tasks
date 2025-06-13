package domain

import "fmt"

type Task struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t *Task) Validate() error {
	if len(t.Title) == 0 {
		return fmt.Errorf("title is required")
	}
	if len(t.Title) < 3 {
		return fmt.Errorf("title must be at least 3 characters")
	}
	return nil
}
