package api

type RestResponse struct {
	Clan Clan `json:"clan"`
}
type Clan struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}
