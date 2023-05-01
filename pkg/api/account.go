package api

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// URL data can be found in url.go
// Query strings can be found in query.go
type APIAccount struct {
	ID          int
	Credentials Credentials
	Logger      *HTTPLogger
	Token       string
	Mutex       *sync.Mutex
	WG          *sync.WaitGroup
	CurrentKey  int
	CurrentSize int
	Keys        []Key
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Key struct {
	ID          string   `json:"id"`
	Developerid string   `json:"developerId"`
	Tier        string   `json:"tier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Origins     any      `json:"origins"`
	Scopes      []string `json:"scopes"`
	Cidrranges  []string `json:"cidrRanges"`
	ValidUntil  any      `json:"validUntil"`
	Key         string   `json:"key"`
}

func (a *APIAccount) login(client *APIClient) error {
	// Sending http request
	credBody, err := json.Marshal(a.Credentials)
	if err != nil {
		return err
	}
	resp, err := client.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(credBody).
		Post(BaseURL + LoginEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   fmt.Sprintf("APIAccount#%d", a.ID),
		Method:   "GET",
		Endpoint: LoginEndpoint,
	}
	client.HTTPLogger.Do(hf, resp)
	logresp := new(LoginResponse)
	err = json.Unmarshal(resp.Body(), &logresp)
	if err != nil {
		return err
	}
	a.Mutex.Lock()
	a.Token = logresp.TemporaryAPIToken
	a.Mutex.Unlock()
	return nil
}
func (a *APIAccount) getKeys(client *APIClient) error {
	// Sending http request
	resp, err := client.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(a.Token).
		Post(BaseURL + KeyListEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   fmt.Sprintf("APIAccount#%d", a.ID),
		Method:   "POST",
		Endpoint: KeyListEndpoint,
	}
	client.HTTPLogger.Do(hf, resp)
	keyresp := new(KeyResponse)
	err = json.Unmarshal(resp.Body(), &keyresp)
	if err != nil {
		return err
	}
	a.Mutex.Lock()
	a.Keys = keyresp.Keys
	a.CurrentSize = len(a.Keys)
	a.Mutex.Unlock()
	return nil
}
func (a *APIAccount) createKey(client *APIClient) error {
	keyindex := a.CurrentKey
	token := a.Token
	a.CurrentKey++
	// Sending http request
	key := new(Key)
	key.Name = fmt.Sprintf("KEY_%d", keyindex)
	key.Description = fmt.Sprintf("Created on %s", time.Now().Format(time.UnixDate))
	key.Cidrranges = []string{client.IP}
	a.Keys = append(a.Keys, *key)
	keyBody, err := json.Marshal(key)
	if err != nil {
		return err
	}
	resp, err := client.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		SetBody(keyBody).
		Post(BaseURL + CreateKeyEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   fmt.Sprintf("APIAccount#%d", a.ID),
		Method:   "POST",
		Endpoint: CreateKeyEndpoint,
	}
	client.HTTPLogger.Do(hf, resp)
	keyresp := new(KeyResponse)
	err = json.Unmarshal(resp.Body(), &keyresp)
	if err != nil {
		return err
	}
	return nil
}
func (a *APIAccount) revokeKey(client *APIClient) error {
	key := a.Keys[a.CurrentKey]
	a.CurrentKey--
	// Sending http request
	revokeBody, err := json.Marshal(key)
	if err != nil {
		return err
	}
	resp, err := client.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(a.Token).
		SetBody(revokeBody).
		Post(BaseURL + RevokeKeyEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   fmt.Sprintf("APIAccount#%d", a.ID),
		Method:   "POST",
		Endpoint: RevokeKeyEndpoint,
	}
	client.HTTPLogger.Do(hf, resp)
	keyresp := new(KeyResponse)
	err = json.Unmarshal(resp.Body(), &keyresp)
	if err != nil {
		return err
	}
	return nil
}
func (a *APIAccount) FillKeys(client *APIClient) error {
	a.CurrentSize = len(a.Keys)
	a.CurrentKey = a.CurrentSize
	for i := a.CurrentSize; i < 10; i++ {
		a.createKey(client)
	}
	return nil
}
func (a *APIAccount) SanitizeKeys(client *APIClient) error {
	a.CurrentSize = len(a.Keys)
	a.CurrentKey = a.CurrentSize - 1
	for i := 0; i < a.CurrentSize; i++ {
		a.revokeKey(client)
	}
	return nil
}
