package api

import (
	"bytes"
	"fmt"
	"net/http"
)

// URL data can be found in url.go
type APIAccount struct {
	Credentials Credentials
	Keys        []string
}
type Credentials struct {
	Email    string
	Password string
}

func (a *APIAccount) Login(client *http.Client) error {
	URL := BaseUrl + LoginEndpoint
	// creating http.Request
	body := bytes.NewReader([]byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, a.Credentials.Email, a.Credentials.Password)))
	req, err := http.NewRequest(http.MethodPost, URL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// handle http.Response
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
	return nil
}
