package model

import "errors"

var (
	ErrTaskTitleEmpty = errors.New("title is empty")
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return ErrTaskTitleEmpty
	}
	return nil
}
