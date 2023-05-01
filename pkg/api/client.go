package api

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"github.com/vexulansis/COCParser/internal/storage"
)

type APIClient struct {
	Client      *resty.Client
	Logger      *logrus.Logger
	HTTPLogger  *HTTPLogger
	Mutex       *sync.Mutex
	WG          *sync.WaitGroup
	IP          string
	Accounts    []*APIAccount
	CurrentAcc  int
	CurrentSize int
}

func NewClient(DBClient *storage.DBClient) (*APIClient, error) {
	client := new(APIClient)
	// Default resty client
	client.Client = resty.New()
	// Custom loggers
	client.HTTPLogger = initHTTPLogger()
	// Filling IP field
	err := client.getIP()
	if err != nil {
		return nil, err
	}
	// Creating Mutex
	client.Mutex = &sync.Mutex{}
	// Creating WaitGroup
	client.WG = &sync.WaitGroup{}
	// Getting accounts from DB
	client.Accounts = convertAccounts(DBClient.APIAccounts)
	client.CurrentSize = len(client.Accounts)
	// Getting keys from COC API
	err = client.GetKeys()
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
	// Logging http response
	hf := HTTPFields{
		Source:   "APIClient",
		Method:   "POST",
		Endpoint: IPURL,
	}
	c.HTTPLogger.Do(hf, resp)
	c.IP = string(resp.Body())
	return nil
}
func (c *APIClient) GetKeys() error {
	c.WG.Add(c.CurrentSize)
	c.CurrentAcc = c.CurrentSize - 1
	for i := 0; i < c.CurrentSize; i++ {
		c.Mutex.Lock()
		acc := c.CurrentAcc
		c.CurrentAcc--
		c.Mutex.Unlock()
		go func() {
			if c.Accounts[acc].login(c) == nil {
				if err := c.Accounts[acc].getKeys(c); err != nil {
					c.Logger.Error(err)
				}
			}
			c.WG.Done()
		}()
	}
	c.WG.Wait()
	return nil
}

func (c *APIClient) FillKeys() error {
	for _, acc := range c.Accounts {
		if err := acc.login(c); err != nil {
			return err
		}
		if err := acc.FillKeys(c); err != nil {
			return err
		}
	}
	return nil
}
func (c *APIClient) SanitizeKeys() error {
	for _, acc := range c.Accounts {
		if err := acc.login(c); err != nil {
			return err
		}
		if err := acc.SanitizeKeys(c); err != nil {
			return err
		}
	}
	return nil
}
func (c *APIClient) CreateKeyPool() []string {
	keys := []string{}
	for _, a := range c.Accounts {
		for _, k := range a.Keys {
			keys = append(keys, k.Key)
		}
	}
	return keys
}
func convertAccounts(accounts []*storage.APIAccount) []*APIAccount {
	conv := []*APIAccount{}
	for _, a := range accounts {
		c := new(APIAccount)
		c.ID = a.ID
		c.Credentials = Credentials(a.Credentials)
		c.WG = &sync.WaitGroup{}
		c.Mutex = &sync.Mutex{}
		conv = append(conv, c)
	}
	return conv
}
