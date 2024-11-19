package usecase

import (
	"github.com/clean_architecture_beta/model"
	"github.com/clean_architecture_beta/repository"
)

type TaskUsecase interface {
	CreateTask(task string) (int, error)
	GetTask(id int) (*model.Task, error)
	UpdateTask(id int, titile string) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	tr repository.TaskRepositoruy
}

func NewTaskUsecase(tr repository.TaskRepositoruy) TaskUsecase {
	return &taskUsecase{tr: tr}
}

func (tu *taskUsecase) CreateTask(title string) (int, error) {
	task := model.Task{Title: title}
	// titleが空文字の場合はエラーを返す
	if err := task.Validate(); err != nil {
		return 0, err
	}

	id, err := tu.tr.Create(&task)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (tu *taskUsecase) GetTask(id int) (*model.Task, error) {
	task, err := tu.tr.Read(id)
	if err != nil {
		return nil, err
	}
	return task, err
}

func (tu *taskUsecase) UpdateTask(id int, title string) error {
	task := model.Task{ID: id, Title: title}
	return tu.tr.Update(&task)
}

func (tu *taskUsecase) DeleteTask(id int) error {
	return tu.tr.Delete(id)
}
