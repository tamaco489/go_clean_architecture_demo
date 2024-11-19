package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/clean_architecture_beta/model"
	"github.com/clean_architecture_beta/usecase"
	"github.com/labstack/echo/v4"
)

type TaskController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
}

type taskController struct {
	tu usecase.TaskUsecase
}

func NewTaskController(tu usecase.TaskUsecase) TaskController {
	return &taskController{tu: tu}
}

func (tc *taskController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error from task controller | %v", err.Error())
		return c.JSON(http.StatusBadRequest, msg.Error())
	}
	task, err := tc.tu.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}

func (tc *taskController) Create(c echo.Context) error {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	createdID, err := tc.tu.CreateTask(task.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// TODO: 201 Createdでモデルそのまま返却したい
	return c.JSON(http.StatusCreated, createdID)
}
