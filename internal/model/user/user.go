package user

type UserID int

type UserName string

type User struct {
	ID   UserID
	Name UserName
}

func (u *User) HasID() bool {
	return u.ID != 0
}
