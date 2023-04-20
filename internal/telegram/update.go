package tg

type RestResponse struct {
	Result []Update `json:"result"`
}

// https://core.telegram.org/bots/api#update
type Update struct {
	Update_id int     `json:"update_id"`
	Message   Message `json:"message"`
}
