package repository

import (
	"database/sql"
	"go-sample/model"
)

type TaskRepository interface {
	Create(task *model.Task) (int, error)
	Read(id int) (*model.Task, error)
	Reads() ([]model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

type taskRepositoryImpl struct{
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepositoryImpl {
	return &taskRepositoryImpl{db: db}
}

func (r *taskRepositoryImpl) Create(task *model.Task) (int, error){
	stmt := `INSERT INTO tasks (title) VALUES (?) RETURNING id`
	err := r.db.QueryRow(stmt, task.Title).Scan(&task.ID)
	return task.ID, err
}

func (r *taskRepositoryImpl) Read(id int) (*model.Task, error){
	stmt := `SELECT id, title FROM tasks WHERE id = ?`
	task := model.Task{}
	err := r.db.QueryRow(stmt,id).Scan(&task.ID,&task.Title)
	return &task, err
}

func (r *taskRepositoryImpl) Reads() ([]model.Task, error) {
    stmt := "SELECT id, title FROM tasks"
    rows, err := r.db.Query(stmt)
    if err != nil {
        return nil, err
    }

    var tasks []model.Task
    for rows.Next() {
        var task model.Task
        err := rows.Scan(&task.ID, &task.Title)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

func (r *taskRepositoryImpl) Update(task *model.Task) error{
	stmt := "UPDATE tasks SET title = ? WHERE id = ?"
	rows,err := r.db.Exec(stmt, task.Title, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil{
		return err
	}
	if rowsAffected == 0{
		return sql.ErrNoRows
	}

return nil
}

func (r *taskRepositoryImpl) Delete(id int ) error{
	stmt := "DELETE FROM tasks WHERE id = ?"
	rows, err := r.db.Exec(stmt, id);
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil{
		return err
	}
	if rowsAffected == 0{
		return sql.ErrNoRows
	}

return nil
}