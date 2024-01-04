package area

import "errors"

var (
    ErrUserIsNotExists = errors.New("user is not exist")
	ErrAreaIsAlreadyExists = errors.New("default area is already exist")
)
