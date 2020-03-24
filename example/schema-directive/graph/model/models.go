package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role Role
}

func (u *User) HasRole(role Role) bool {
	return u.Role == role || u.Role == RoleAdmin
}
