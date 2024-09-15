package auth

import "errors"

type Role int

const (
	RoleUser Role = iota
	RoleAdmin
)

func (r Role) String() string {
	return [...]string{"user", "admin"}[r]
}

var stringToRoleMapper = map[string]Role{
	"user":  RoleUser,
	"admin": RoleAdmin,
}

func RoleFromString(s string) (Role, error) {
	if role, exists := stringToRoleMapper[s]; exists {
		return role, nil
	}
	return Role(0), errors.New("invalid role")
}
