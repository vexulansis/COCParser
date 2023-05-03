package api

import (
	"github.com/go-resty/resty/v2"
)

type APIClient struct {
	Logger    *APILogger
	Client    *resty.Client
	IP        string
	TagChan   chan string
	ClanChan  chan *Clan
	ErrorChan chan error
}

func NewClient() (*APIClient, error) {
	// Creating AC example
	apiClient := &APIClient{
		Client: resty.New(),
	}
	// Initializing logger
	apiClient.Logger = initAPILogger()
	// Getting IP
	err := apiClient.getIP()
	if err != nil {
		return nil, err
	}
	// Creating tag channel
	apiClient.TagChan = make(chan string)
	// Creating clan channel
	apiClient.ClanChan = make(chan *Clan)
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
		Source:   "APIClient",
		Method:   "POST",
		Endpoint: IPURL,
	}
	c.Logger.Print(f, resp)
	c.IP = string(resp.Body())
	return nil
}
