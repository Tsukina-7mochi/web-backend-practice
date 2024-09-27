package model

type Todo struct {
	Ref   string `json:"ref"`
	User  string `json:"user"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func NewTodo(ref string, user string, title string, done bool) *Todo {
	return &Todo{
		Ref:   ref,
		User:  user,
		Title: title,
		Done:  done,
	}
}
