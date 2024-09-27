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

	user := entity.User{
		ID: id,
	}
	if err := row.Scan(&user.Name, &user.DisplayName); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByName(name string) (*entity.User, error) {
	row := r.db.QueryRow(`SELECT id, display_name FROM users WHERE name = $1;`, name)

	user := entity.User{
		Name: name,
	}
	if err := row.Scan(&user.ID, &user.DisplayName); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(id uint) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}
