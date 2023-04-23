package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/joho/godotenv"
)

var apiURLprefix = "https://api.clashofclans.com/v1/clans/%23"

func defaultRequest(URL string, token string) (*http.Request, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	err = godotenv.Load()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")
	return req, nil
}
func GetClanByTag(clanTag string, token string) (*Clan, error) {
	reqURL := apiURLprefix + clanTag
	req, err := defaultRequest(reqURL, token)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		clan := new(Clan)
		err = json.Unmarshal(body, &clan)
		return clan, nil
	} else {
		clientErr := new(ClientError)
		err = json.Unmarshal(body, &clientErr)
	}
	return nil, nil
}
