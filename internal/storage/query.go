package storage

const (
	getAccountsQuery = "select * from Auth"
	addAccountQuery  = "insert into Auth(email,password) values ($1,$2)"
)
