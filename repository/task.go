package repository

import (
	"database/sql"

	"github.com/clean_architecture_beta/model"
)

type TaskRepositoruy interface {
	Create(task *model.Task) (int, error)
	Read(id int) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

// このパッケージ内でしか使用しないためローカルで定義
type taskRepositoryImpl struct {
	db *sql.DB
}

// こちらは外部から呼び出すためにグローバルで定義
func NewTaskRepository(db *sql.DB) *taskRepositoryImpl {
	return &taskRepositoryImpl{db: db}
}

func (tr *taskRepositoryImpl) Create(task *model.Task) (int, error) {
	stmt := `INSERT INTO tasks(title) VALUES(?) RETURNING id`
	if err := tr.db.QueryRow(stmt, task.Title).Scan(&task.ID); err != nil {
		return 0, err
	}
	return task.ID, nil
}

func (tr *taskRepositoryImpl) Read(id int) (*model.Task, error) {
	stmt := `SELECT id, title FROM tasks WHERE id = ?`
	task := &model.Task{}
	if err := tr.db.QueryRow(stmt, id).Scan(&task.ID, &task.Title); err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *taskRepositoryImpl) Update(task *model.Task) error {
	stmt := `UPDATE tasks SET title = ? WHERE id = ?`
	rows, err := tr.db.Exec(stmt, task.Title, task.ID)
	if err != nil {
		return err
	}

	// RowsAffected: 更新、挿入、または削除によって影響を受ける行の数を返す
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (tr *taskRepositoryImpl) Delete(id int) error {
	stmt := `DELETE FROM tasks WHERE id = ?`
	rowsAffected, err := tr.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	if rowsAffected == nil {
		return sql.ErrNoRows
	}
	return nil
}
