package api

type ClientError struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
