package model

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func NewUser(name string, displayName string) *User {
	return &User{
		Name:        name,
		DisplayName: displayName,
	}
}
