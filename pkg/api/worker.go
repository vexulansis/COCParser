package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Worker struct {
	ID    int
	Token string
	Pool  *Pool
	Quit  chan bool
}

func NewWorker(ID int, pool *Pool) *Worker {
	return &Worker{
		ID:   ID,
		Pool: pool,
		Quit: make(chan bool),
	}
}
func (w *Worker) Start() {
	// f := APILoggerFields{
	// 	Source:      fmt.Sprintf("APIWORKER#%d", w.ID),
	// 	Method:      "START",
	// 	Subject:     "<---",
	// 	Destination: "TASKCHANNEL",
	// }
	// w.Pool.AC.Logger.Print(f, 0)

	for {
		select {
		case task := <-w.Pool.AC.TaskChan:
			w.Process(task)
		case <-w.Quit:
			return
		}
	}
}
func (w *Worker) Process(task *Task) error {
	defer w.Pool.WG.Done()
	switch t := task.Data.(type) {
	case *Account:
		w.login(t)
		w.getKeys(t)
	case string:
		w.getClanByTag(t)
	}
	return nil
}
func (w *Worker) Stop() {
	f := APILoggerFields{
		Source:      fmt.Sprintf("APIWORKER#%d", w.ID),
		Method:      "STOP",
		Subject:     "<---",
		Destination: "TASKCHANNEL",
	}
	w.Pool.AC.Logger.Print(f, 0)
	go func() {
		w.Quit <- true
	}()
}
func (w *Worker) login(a *Account) error {
	// Sending http request
	body, err := json.Marshal(a)
	if err != nil {
		return err
	}
	resp, err := w.Pool.AC.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(BaseURL + LoginEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	f := APILoggerFields{
		Source:      fmt.Sprintf("APIWORKER#%d", w.ID),
		Method:      "POST",
		Subject:     a.Email,
		Destination: LoginEndpoint,
	}
	w.Pool.AC.Logger.Print(f, resp)
	logresp := &LoginResponse{}
	err = json.Unmarshal(resp.Body(), &logresp)
	if err != nil {
		return err
	}
	// Mutex lock for stability
	w.Pool.AC.Mutex.Lock()
	a.Token = logresp.TemporaryAPIToken
	w.Pool.AC.Mutex.Unlock()
	return nil

}
func (w *Worker) getKeys(a *Account) error {
	// Sending http request
	resp, err := w.Pool.AC.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(a.Token).
		Post(BaseURL + KeyListEndpoint)
	if err != nil {
		return err
	}
	// Logging http response
	f := APILoggerFields{
		Source:      fmt.Sprintf("APIWORKER#%d", w.ID),
		Method:      "POST",
		Subject:     a.Email,
		Destination: KeyListEndpoint,
	}
	w.Pool.AC.Logger.Print(f, resp)
	keyresp := &KeyResponse{}
	err = json.Unmarshal(resp.Body(), &keyresp)
	if err != nil {
		return err
	}
	w.Pool.AC.Mutex.Lock()
	a.Keys = keyresp.Keys
	w.Pool.AC.Mutex.Unlock()
	return nil
}
func (w *Worker) getClanByTag(tag string) (*Clan, error) {
	// Making http request
	resp, err := w.Pool.AC.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(w.Token).
		Get(APIURL + ClansEndpoint + "/%23" + tag)
	if err != nil {
		return nil, err
	}
	// Logging http response
	f := APILoggerFields{
		Source:      fmt.Sprintf("APIWORKER#%d", w.ID),
		Method:      "POST",
		Subject:     "#" + tag,
		Destination: ClansEndpoint,
	}
	w.Pool.AC.Logger.Print(f, resp)
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
