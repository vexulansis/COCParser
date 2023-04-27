package api

import (
	"fmt"
	"net/http"
)

type APIClient struct {
	Client   http.Client
	Accounts []APIAccount
}

func NewClient() *APIClient {
	client := new(APIClient)
	client.Client = *http.DefaultClient
	client.Accounts = AuthAccounts()
	for _, acc := range client.Accounts {
		err := acc.Login(&client.Client)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
	return client
}
