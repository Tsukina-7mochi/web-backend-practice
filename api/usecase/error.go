package usecase

import "errors"

var NoUserError error = errors.New("No such user")
var NoTodoError error = errors.New("No such toodo")
