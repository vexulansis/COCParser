package storage

type APIAccount struct {
	ID          int
	Credentials Credentials
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
