package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// URL data can be found in url.go
// Query strings can be found in query.go
type APIAccount struct {
	Credentials Credentials
	Response    LoginResponse
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
	rows, err := db.Query(getAccountsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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
func (a *APIAccount) login(client *resty.Client) error {
	credBody, err := json.Marshal(a.Credentials)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(credBody).
		Post(BaseUrl + LoginEndpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("%d\n", resp.StatusCode())
	}
	err = json.Unmarshal(resp.Body(), &a.Response)
	if err != nil {
		return err
	}
	return nil
}
func (a *APIAccount) getKeys(client *resty.Client) error {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(a.Response.TemporaryAPIToken).
		Post(BaseUrl + KeyListEndpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("%d\n", resp.StatusCode())
	}
	keyresp := new(KeyResponse)
	err = json.Unmarshal(resp.Body(), &keyresp)
	if err != nil {
		return err
	}
	a.Keys = keyresp.Keys
	return nil
}
