package domain

type Role int

const (
	RoleUser Role = iota
	RoleModerator
	RoleSuperModerator
	RoleAdmin
)

type TestInput struct {
	ID   string
	Role Role
}

type TestType struct {
	ID   string
	Role Role
}
