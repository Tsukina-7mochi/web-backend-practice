package entity

type Todo struct {
	ID     uint
	Ref    string
	UserID uint
	Title  string
	Done   bool
}

func NewTodo(id uint, ref string, userID uint, title string, done bool) *Todo {
	return &Todo{
		ID:     id,
		Ref:    ref,
		UserID: userID,
		Title:  title,
		Done:   done,
	}
}
