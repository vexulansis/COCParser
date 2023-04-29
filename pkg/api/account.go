package api

import (
	"database/sql"
	"encoding/json"
)

// URL data can be found in url.go
// Query strings can be found in query.go
type APIAccount struct {
	Credentials Credentials
	Token       string
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

func getAccounts(db *sql.DB) ([]*APIAccount, error) {
	accounts := []*APIAccount{}
	// Rows: email, password
	rows, err := db.Query(getAccountsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Iterating through rows, creating new accounts using credentials
	for rows.Next() {
		cred := Credentials{}
		err := rows.Scan(&cred.Email, &cred.Password)
		if err != nil {
			return nil, err
		}
		acc := &APIAccount{
			Credentials: cred,
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}
func (a *APIAccount) login(client *APIClient) error {
	// Sending http request
	credBody, err := json.Marshal(a.Credentials)
	resp, err := client.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(credBody).
		Post(BaseURL + LoginEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   "APIClient",
		Method:   "GET",
		Endpoint: LoginEndpoint,
	}
	client.HTTPLogger.Do(hf, resp)
	logresp := new(LoginResponse)
	err = json.Unmarshal(resp.Body(), &logresp)
	if err != nil {
		return err
	}
	a.Token = logresp.TemporaryAPIToken
	return nil
}
func (a *APIAccount) getKeys(client *APIClient) error {
	resp, err := client.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(a.Token).
		Post(BaseURL + KeyListEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   "APIClient",
		Method:   "POST",
		Endpoint: KeyListEndpoint,
	}
	client.HTTPLogger.Do(hf, resp)
	keyresp := new(KeyResponse)
	err = json.Unmarshal(resp.Body(), &keyresp)
	if err != nil {
		return err
	}
	a.Keys = keyresp.Keys
	return nil
}
