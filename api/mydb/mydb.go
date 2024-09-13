package mydb

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
}

func (user User) JSON() string {
	return fmt.Sprintf(
		`{"id":%d,"name":"%s"}`,
		user.ID,
		user.Name,
	)
}

type Todo struct {
	ID     int
	UserID int
	Title  string
	Done   bool
}

func (todo Todo) JSON() string {
	return fmt.Sprintf(
		`{"id":%d,"user_id":%d,"title":"%s","done":%t}`,
		todo.ID,
		todo.UserID,
		todo.Title,
		todo.Done,
	)
}

type MyDB struct {
	db *sql.DB
}

type DBInit struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Open(init DBInit) (*MyDB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			init.Host,
			init.Port,
			init.User,
			init.Password,
			init.Name,
		),
	)

	return &MyDB{db: db}, err
}

func (myDB *MyDB) Init() error {
	db := myDB.db

	// Crete users tables
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL NOT NULL,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        UNIQUE (name)
    );`)
	if err != nil {
		return err
	}

	// Create todos table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
        id SERIAL NOT NULL,
        user_id INT NOT NULL,
        title VARCHAR(255) NOT NULL,
        done BOOLEAN NOT NULL DEFAULT FALSE,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`)
	if err != nil {
		return err
	}

	return nil
}

func (myDB *MyDB) AddUser(name string) (int, error) {
	db := myDB.db

	row := db.QueryRow(`INSERT INTO users (name) VALUES ($1) RETURNING id;`, name)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (myDB *MyDB) GetUser(name string) (*User, error) {
	db := myDB.db
	row := db.QueryRow(`SELECT id, name FROM users WHERE name = $1;`, name)

	var id int
	var n string
	err := row.Scan(&id, &n)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:   id,
		Name: n,
	}, nil
}

func (myDB *MyDB) DeleteUser(id int) error {
	db := myDB.db

	_, err := db.Exec(`DELETE FROM todos WHERE user_id = $1;`, id)
	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}

func (myDB *MyDB) AddTodo(userID int, title string) (int, error) {
	db := myDB.db

	row := db.QueryRow(`INSERT INTO todos (user_id, title) VALUES ($1, $2) RETURNING id;`, userID, title)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (myDB *MyDB) ListByUser(userID int) ([]Todo, error) {
	db := myDB.db
	rows, err := db.Query(`SELECT id, user_id, title, done FROM todos WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, err
	}

	todos := make([]Todo, 0)
	for rows.Next() {
		var id int
		var userID int
		var title string
		var done bool

		err = rows.Scan(&id, &userID, &title, &done)
		if err != nil {
			if err == sql.ErrNoRows {
				return []Todo{}, nil
			}
			return nil, err
		}

		todos = append(todos, Todo{
			ID:     id,
			UserID: userID,
			Title:  title,
			Done:   done,
		})
	}

	return todos, nil
}

func (myDB *MyDB) PatchTodo(todoID int, done bool) error {
	db := myDB.db

	_, err := db.Exec(`UPDATE todos SET done = $1 WHERE id = $2;`, done, todoID)

	return err
}

func (myDB *MyDB) BulkDeleteTodos(ids []int) error {
	db := myDB.db

	_, err := db.Exec(`DELETE FROM todos WHERE id = ANY($1);`, pq.Array(ids))

	return err
}
