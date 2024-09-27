package usecase

import (
	"database/sql"
	"main/model"
	. "main/repository"
)

type TodoUsecase struct {
	userRepo UserRepository
	todoRepo TodoRepository
}

func NewTodoUsecase(userRepo UserRepository, todoRepo TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		userRepo: userRepo,
		todoRepo: todoRepo,
	}
}

func (u *TodoUsecase) Create(userName string, title string) error {
	user, err := u.userRepo.GetByName(userName)
	if err == sql.ErrNoRows {
		return NoUserError
	} else if err != nil {
		return err
	}

	if _, err := u.todoRepo.Create(user.ID, title); err != nil {
		return err
	}

	return nil
}

func (u *TodoUsecase) ListByUser(userName string) ([]model.Todo, error) {
	user, err := u.userRepo.GetByName(userName)
	if err == sql.ErrNoRows {
		return nil, NoUserError
	} else if err != nil {
		return nil, err
	}

	todoEntities, err := u.todoRepo.ListByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	todos := make([]model.Todo, 0, len(todoEntities))
	for _, todoEntity := range todoEntities {
		todos = append(todos, *model.NewTodoFromEntity(&todoEntity, userName))
	}

	return todos, nil
}

func (u *TodoUsecase) UpdateDone(userName string, ref string, done bool) error {
	user, err := u.userRepo.GetByName(userName)
	if err == sql.ErrNoRows {
		return NoUserError
	} else if err != nil {
		return err
	}

	todo, err := u.todoRepo.GetByRef(user.ID, ref)
	if err == sql.ErrNoRows {
		return NoTodoError
	} else if err != nil {
		return err
	}

	if err := u.todoRepo.UpdateDone(todo.ID, done); err != nil {
		return err
	}

	return nil
}

func (u *TodoUsecase) Delete(userName string, ref string) error {
	user, err := u.userRepo.GetByName(userName)
	if err == sql.ErrNoRows {
		return NoUserError
	} else if err != nil {
		return err
	}

	todo, err := u.todoRepo.GetByRef(user.ID, ref)
	if err == sql.ErrNoRows {
		return NoTodoError
	} else if err != nil {
		return err
	}

	if err := u.todoRepo.Delete(todo.ID); err != nil {
		return err
	}

	return nil
}
