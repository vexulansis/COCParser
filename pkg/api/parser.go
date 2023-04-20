package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var apiURLprefix = "https://api.clashofclans.com/v1/clans/%23"

func defaultRequest(URL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	err = godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("API_TOKEN")
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	return req, nil
}
func GetClanByTag(clanTag string) (*Clan, error) {
	reqURL := apiURLprefix + clanTag
	req, err := defaultRequest(reqURL)
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
	clan := new(Clan)
	err = json.Unmarshal(body, &clan)
	return clan, nil

}
