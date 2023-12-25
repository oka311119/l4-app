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
	userId string,
	name string,
	createdAt time.Time,
) *Area {
	return &Area{
		ID: id,
		UserID: userId,
		Name: name,
		CreatedAt: createdAt,
	}
}
