package api

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	db "github.com/vexulansis/COCParser/internal/storage"
)

type APIClient struct {
	Client     *resty.Client
	Logger     *logrus.Logger
	HTTPLogger *HTTPLogger
	Accounts   []*APIAccount
	IP         string
}

func NewClient(storage *db.Storage) (*APIClient, error) {
	client := new(APIClient)
	// Default resty client
	client.Client = resty.New()
	// Custom loggers
	client.Logger = initLogger()
	client.HTTPLogger = initHTTPLogger()
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
	// Logging http response
	hf := HTTPFields{
		Source:   "APIClient",
		Method:   "POST",
		Endpoint: KeyListEndpoint,
	}
	c.HTTPLogger.Do(hf, resp)
	c.IP = string(resp.Body())
	return nil
}
func (c *APIClient) getKeys() error {
	for _, a := range c.Accounts {
		err := a.login(c)
		if err != nil {
			return err
		}
		err = a.getKeys(c)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *APIClient) GetClanByTag(tag string, key string) error {
	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(key).
		Get(APIURL + ClansEndpoint + "/%23" + tag)
	if err != nil {
		return err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   "APIClient",
		Method:   "GET",
		Endpoint: ClansEndpoint,
	}
	c.HTTPLogger.Do(hf, resp)
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
