package entities

type User struct {
	Id       int32 `pg:",pk,unique"`
	Username string
	UserType string
	Email    string
	Password string
}
