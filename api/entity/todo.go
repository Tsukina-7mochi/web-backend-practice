package entity

type Todo struct {
	ID     uint
	Ref    string
	UserID uint
	Title  string
	Done   bool
}
