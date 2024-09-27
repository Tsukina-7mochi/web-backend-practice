package usecase

import (
	"database/sql"
	"main/model"
	. "main/repository"
)

type UserUsecase struct {
	userRepo UserRepository
}

func NewUserUsecase(userRepo UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Create(name string, displayName string) error {
	if _, err := u.userRepo.Create(name, displayName); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Get(name string) (*model.User, error) {
	user, err := u.userRepo.GetByName(name)
	if err == sql.ErrNoRows {
		return nil, NoUserError
	} else if err != nil {
		return nil, err
	}

	return model.NewUser(user.Name, user.DisplayName), nil
}

func (u *UserUsecase) Delete(name string) error {
	user, err := u.userRepo.GetByName(name)
	if err != nil {
		return err
	}

	if err := u.userRepo.Delete(user.ID); err != nil {
		return err
	}

	return nil
}
