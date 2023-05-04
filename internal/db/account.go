package db

type Account struct {
	ID          int
	Credentials Credentials
}
type Credentials struct {
	Email    string
	Password string
}
