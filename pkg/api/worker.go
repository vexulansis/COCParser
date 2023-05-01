package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	kafka "github.com/segmentio/kafka-go"
)

type Worker struct {
	ID       int
	Key      string
	Client   *resty.Client
	Logger   *HTTPLogger
	Mutex    *sync.Mutex
	TagPool  *kafka.Reader
	ClanPool *kafka.Writer
	WG       *sync.WaitGroup
	Status   chan int
	Output   chan error
}
type TagChunk struct {
	Tags []string
}
type ClanChunk struct {
	Clans []*Clan
}

func NewWorker(key string, id int) *Worker {
	w := &Worker{
		ID:     id,
		Key:    key,
		Client: resty.New(),
		Logger: initHTTPLogger(),
		Mutex:  new(sync.Mutex),
		WG:     new(sync.WaitGroup),
		Status: make(chan int),
		Output: make(chan error),
	}
	return w
}

// pseudo
func (w *Worker) Execute() {
	defer w.WG.Done()
	var tagChunk *TagChunk
	var clanChunk *ClanChunk
	// Getting tags from TagPool
	tagMsg, err := w.TagPool.ReadMessage(context.Background())
	if err != nil {
		w.Output <- err
	}
	// Generating tag array
	err = json.Unmarshal(tagMsg.Value, &tagChunk)
	if err != nil {
		w.Output <- err
	}
	// Iterating through array till 404
	for _, tag := range tagChunk.Tags {
		clan := w.GetClanByTag(tag)
		// Prevent rate limit
		time.Sleep(time.Millisecond * 15)
		clanChunk.Clans = append(clanChunk.Clans, clan)
	}
	// Marshalling clans
	clanMsg, err := json.Marshal(clanChunk)
	if err != nil {
		w.Output <- err
	}
	// Sending clans to ClanPool
	w.ClanPool.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(strconv.Itoa(w.ID)),
			Value: clanMsg,
		})
}
func (w *Worker) GetClanByTag(tag string) *Clan {
	// Making http request
	resp, err := w.Client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(w.Key).
		Get(APIURL + ClansEndpoint + "/%23" + tag)
	if err != nil {
		w.Output <- err
	}
	// Logging http response
	hf := HTTPFields{
		Source:   fmt.Sprintf("Worker#%d", w.ID),
		Method:   "GET",
		Endpoint: ClansEndpoint,
	}
	w.Logger.Do(hf, resp)
	w.Status <- resp.StatusCode()
	switch resp.StatusCode() {
	case http.StatusOK:
		clan := new(Clan)
		json.Unmarshal(resp.Body(), &clan)
		return clan
	}
	return nil
}
