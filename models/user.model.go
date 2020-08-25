package models

type User struct {
	Id       int
	Username string
	Password string
	RoleId   int
	Status   int
}
