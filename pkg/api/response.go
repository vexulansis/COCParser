package api

type LoginResponse struct {
	Status            Status `json:"status"`
	TemporaryAPIToken string `json:"temporaryAPIToken"`
}

type KeyResponse struct {
	Status Status `json:"status"`
	Keys   []Key  `json:"keys,omitempty"`
}
type Status struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  any    `json:"detail"`
}
