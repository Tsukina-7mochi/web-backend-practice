package repository

import (
	"database/sql"
	"main/entity"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(userID uint, title string) (int, error) {
	row := r.db.QueryRow(`INSERT INTO todos (user_id, title) VALUES ($1, $2) RETURNING id;`, userID, title)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}

func (r *TodoRepository) GetByRef(userID uint, ref string) (*entity.Todo, error) {
	row := r.db.QueryRow(`SELECT id, user_id, title, done FROM todos WHERE user_id = $1 AND ref = $2;`, userID, ref)

	var id uint
	var title string
	var done bool
	if err := row.Scan(&id, &userID, &title, &done); err != nil {
		return nil, err
	}

	todo := entity.NewTodo(id, ref, userID, title, done)
	return todo, nil
}

func (r *TodoRepository) ListByUserID(userID uint) ([]entity.Todo, error) {
	rows, err := r.db.Query(`SELECT id, ref, user_id, title, done FROM todos WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, err
	}

	todos := make([]entity.Todo, 0)
	for rows.Next() {
		var id uint
		var ref string
		var userID uint
		var title string
		var done bool

		if err = rows.Scan(&id, &ref, &userID, &title, &done); err != nil {
			return nil, err
		}

		todo := entity.NewTodo(id, ref, userID, title, done)
		todos = append(todos, *todo)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	return todos, nil

}

func (r *TodoRepository) UpdateDone(id uint, done bool) error {
	_, err := r.db.Exec(`UPDATE todos SET done = $1 WHERE id = $2;`, done, id)
	return err
}

func (r *TodoRepository) Delete(id uint) error {
	_, err := r.db.Exec(`DELETE FROM todos WHERE id = $1;`, id)
	return err
}
