package tg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

var botURLprefix = "https://api.telegram.org/bot"

// https://core.telegram.org/bots/api#getupdates
func (b *TGBot) getUpdates(offset int) ([]Update, error) {
	botURL := botURLprefix + b.Token
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	restResp := new(RestResponse)
	if err = json.Unmarshal(body, &restResp); err != nil {
		return nil, err
	}
	return restResp.Result, nil
}

// https://core.telegram.org/bots/api#sendmessage
func (b *TGBot) sendMessage(msg *BotMessage) (*http.Response, error) {
	botURL := botURLprefix + b.Token
	buf, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
