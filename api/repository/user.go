package repository

import (
	"database/sql"
	"main/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(name string, displayName string) (uint, error) {
	row := r.db.QueryRow(`INSERT INTO users (name, display_name) VALUES ($1, $2) RETURNING id;`, name, displayName)

	var id uint
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) Get(id uint) (*entity.User, error) {
	row := r.db.QueryRow(`SELECT name, display_name FROM users WHERE id = $1;`, id)

	var name string
	var displayName string
	if err := row.Scan(&name, &displayName); err != nil {
		return nil, err
	}

	return entity.NewUser(id, name, displayName), nil
}

func (r *UserRepository) GetByName(name string) (*entity.User, error) {
	row := r.db.QueryRow(`SELECT id, display_name FROM users WHERE name = $1;`, name)

	var id uint
	var displayName string
	if err := row.Scan(&id, &displayName); err != nil {
		return nil, err
	}

	return entity.NewUser(id, name, displayName), nil
}

func (r *UserRepository) Delete(id uint) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}
