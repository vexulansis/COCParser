package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Worker struct {
	ID    int
	Token string
	AC    *APIClient
}

func NewWorker(id int, ac *APIClient) *Worker {
	return &Worker{
		ID: id,
		AC: ac,
	}
}
func (w *Worker) getClans(AC *APIClient) {
	for tag := range w.AC.TagChan {
		clan, err := w.getClanByTag(tag, AC)
		if err != nil {
			w.AC.ErrorChan <- err
		}
		w.AC.ClanChan <- clan
	}
}
func (w *Worker) getClanByTag(tag string, AC *APIClient) (*Clan, error) {
	// Making http request
	resp, err := w.AC.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(w.Token).
		Get(APIURL + ClansEndpoint + "/%23" + tag)
	if err != nil {
		return nil, err
	}
	// Logging http response
	f := APILoggerFields{
		Source:   fmt.Sprintf("Worker#%d", w.ID),
		Method:   "POST",
		Endpoint: ClansEndpoint,
	}
	AC.Logger.Print(f, resp)
	if resp.StatusCode() == http.StatusOK {
		clan := &Clan{}
		json.Unmarshal(resp.Body(), &clan)
		if err != nil {
			return nil, err
		}
		return clan, nil
	}
	return nil, nil
}
