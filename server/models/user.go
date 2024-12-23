package models

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}
