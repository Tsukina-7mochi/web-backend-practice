package entity

type User struct {
	ID          uint
	Name        string
	DisplayName string
}

func NewUser(id uint, name string, displayName string) *User {
	return &User{
		ID:          id,
		Name:        name,
		DisplayName: displayName,
	}
}
