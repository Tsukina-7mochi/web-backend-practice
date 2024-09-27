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
	row := r.db.QueryRow(`SELECT id, title, done FROM todos WHERE user_id = $1 AND ref = $2;`, userID, ref)

	todo := entity.Todo{
		UserID: userID,
		Ref:    ref,
	}
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) ListByUserID(userID uint) ([]entity.Todo, error) {
	rows, err := r.db.Query(`SELECT id, ref, user_id, title, done FROM todos WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, err
	}

	todos := make([]entity.Todo, 0)
	for rows.Next() {
		todo := entity.Todo{}
		if err = rows.Scan(&todo.ID, &todo.Ref, &todo.UserID, &todo.Title, &todo.Done); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
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
