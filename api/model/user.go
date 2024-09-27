package model

import "main/entity"

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func NewUserFromEntity(user *entity.User) *User {
	return &User{
		Name:        user.Name,
		DisplayName: user.DisplayName,
	}
}
