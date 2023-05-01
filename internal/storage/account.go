package storage

type APIAccount struct {
	Credentials Credentials
	ID          int
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
