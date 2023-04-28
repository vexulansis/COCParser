package api

type LoginResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int    `json:"sessionExpiresInSeconds"`
	TemporaryAPIToken       string `json:"temporaryAPIToken"`
	SwaggerURL              string `json:"swaggerUrl"`
}

type KeyResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int    `json:"sessionExpiresInSeconds"`
	Keys                    []Key  `json:"keys,omitempty"`
}
type Status struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  any    `json:"detail"`
}
