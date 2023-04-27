package api

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Auth struct {
	Accounts []Credentials
}

func AuthAccounts() []APIAccount {
	auth := new(Auth)
	accounts := []APIAccount{}
	_, err := toml.DecodeFile("auth.toml", &auth)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, cred := range auth.Accounts {
		acc := APIAccount{
			Credentials: cred,
		}
		accounts = append(accounts, acc)
	}
	return accounts
}
