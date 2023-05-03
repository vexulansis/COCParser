package api

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string
	Keys     []Key
}
type Key struct {
	ID          string   `json:"id"`
	Developerid string   `json:"developerId"`
	Tier        string   `json:"tier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Origins     any      `json:"origins"`
	Scopes      []string `json:"scopes"`
	Cidrranges  []string `json:"cidrRanges"`
	ValidUntil  any      `json:"validUntil"`
	Key         string   `json:"key"`
}
