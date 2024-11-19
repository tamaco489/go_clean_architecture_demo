package main

import (
	"database/sql"
	"log"

	"github.com/clean_architecture_beta/controller"
	"github.com/clean_architecture_beta/repository"
	"github.com/clean_architecture_beta/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

var (
	port      = "8080"
	sqlDriver = "sqlite3"
	dbName    = "./test.db"
	queryExec = `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL
		)
	`
)

func initDB() (*sql.DB, error) {
	return sql.Open(sqlDriver, dbName)
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(queryExec)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // panicが発生した場合に回復させるミドルウェア

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"http://localhost:3000"}, // リクエストを許可するオリジンを指定
		AllowMethods: []string{
			echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e.GET("/tasks/:id", taskController.Get)
	e.POST("/tasks", taskController.Create)

	e.Start(":" + port)
}
