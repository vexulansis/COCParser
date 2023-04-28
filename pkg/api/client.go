package api

import (
	"github.com/go-resty/resty/v2"
	db "github.com/vexulansis/COCParser/internal/storage"
)

type APIClient struct {
	Client   *resty.Client
	Accounts []*APIAccount
	IP       string
}

func NewClient(storage *db.Storage) (*APIClient, error) {
	client := new(APIClient)
	client.Client = resty.New()
	// Filling IP field
	err := client.getIP()
	if err != nil {
		return nil, err
	}
	// Getting accounts from DB
	accounts, err := getAccounts(storage.DB)
	if err != nil {
		return nil, err
	}
	client.Accounts = accounts
	// Getting keys from COC API
	err = client.getKeys()
	if err != nil {
		return nil, err
	}
	return client, nil
}
func (c *APIClient) getIP() error {
	resp, err := c.Client.R().Get(IPURL)
	if err != nil {
		return err
	}
	c.IP = string(resp.Body())
	return nil
}
func (c *APIClient) getKeys() error {
	for _, a := range c.Accounts {
		err := a.login(c.Client)
		if err != nil {
			return err
		}
		err = a.getKeys(c.Client)
		if err != nil {
			return err
		}
	}
	return nil
}
