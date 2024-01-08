package entity

type Area struct {
	ID string
	UserID string
	Name string
}

const DefaultAreaName = "$default"

func NewArea(
	id string,
	userID string,
	name string,
) *Area {
	return &Area{
		ID: id,
		UserID: userID,
		Name: name,
    }
}
