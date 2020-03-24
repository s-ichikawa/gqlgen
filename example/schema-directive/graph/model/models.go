package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role Role
}

func (u *User) IsPerson() {

}
func (u *User) IsHuman() {

}

func (u *User) HasRole(role Role) bool {
	switch u.Role {
	case RoleSuperAdmin:
		return true
	case RoleAdmin:
		return role != RoleSuperAdmin
	default:
		return u.Role == role
	}
}
