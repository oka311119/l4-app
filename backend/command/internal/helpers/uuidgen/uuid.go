package uuidgen

import uuid "github.com/satori/go.uuid"

type UUIDGenerator interface {
	V4() string
}

type UUID struct {}

func (*UUID) V4() string {
    return uuid.NewV4().String()
}
