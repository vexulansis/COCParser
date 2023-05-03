package api

import (
	"sync"

	"github.com/go-resty/resty/v2"
)

type APIClient struct {
	Logger      *APILogger
	Client      *resty.Client
	Mutex       *sync.Mutex
	AccountPool []*Account
	KeyPool     []Key
	IP          string
	TaskChan    chan *Task
	ErrorChan   chan error
}

func NewClient() (*APIClient, error) {
	// Creating AC example
	apiClient := &APIClient{
		Client: resty.New(),
		Mutex:  &sync.Mutex{},
	}
	// Initializing logger
	apiClient.Logger = initAPILogger()
	// Getting IP
	err := apiClient.getIP()
	if err != nil {
		return nil, err
	}
	// Creating error channel
	apiClient.TaskChan = make(chan *Task)
	// Creating error channel
	apiClient.ErrorChan = make(chan error)
	return apiClient, nil
}
func (c *APIClient) getIP() error {
	// Creating http request
	resp, err := c.Client.R().Get(IPURL)
	if err != nil {
		return err
	}
	// Logging http response
	f := APILoggerFields{
		Source:      "APICLIENT",
		Method:      "POST",
		Destination: IPURL,
	}
	c.Logger.Print(f, resp)
	c.IP = string(resp.Body())
	return nil
}
func (c *APIClient) CreateKeyPool() {
	for _, a := range c.AccountPool {
		for _, k := range a.Keys {
			c.KeyPool = append(c.KeyPool, k)
		}
	}
	f := APILoggerFields{
		Source:  "APICLIENT",
		Method:  "CREATE",
		Subject: "KEYPOOL",
	}
	c.Logger.Print(f, 0)
}
