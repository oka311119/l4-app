package entity

import "time"

type Area struct {
	ID string
	UserID string
	Name string
	CreatedAt time.Time
}

func NewArea(
	id string,
	userID string,
	name string,
	createdAt time.Time,
) *Area {
	return &Area{
		ID: id,
		UserID: userID,
		Name: name,
		CreatedAt: createdAt,
	}
}
