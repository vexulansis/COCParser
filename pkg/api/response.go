package api

type LoginResponse struct {
	TemporaryAPIToken string `json:"temporaryAPIToken"`
	Status            Status `json:"status"`
}

type KeyResponse struct {
	Status Status `json:"status"`
	Keys   []Key  `json:"keys,omitempty"`
}
type Status struct {
	Detail  any    `json:"detail"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}
