package api

type RestResponse struct {
	Clan Clan `json:"clan"`
}
type Clan struct {
	Description            string `json:"description"`
	Tag                    string `json:"tag"`
	Name                   string `json:"name"`
	Type                   string `json:"type"`
	RequiredTrophies       int    `json:"requiredTrophies"`
	ClanLevel              int    `json:"clanLevel"`
	ClanPoints             int    `json:"clanPoints"`
	ClanVersusPoints       int    `json:"clanVersusPoints"`
	ID                     int
	WarWinStreak           int  `json:"warWinStreak"`
	WarWins                int  `json:"warWins"`
	WarTies                int  `json:"warTies"`
	WarLosses              int  `json:"warLosses"`
	Members                int  `json:"members"`
	RequiredVersusTrophies int  `json:"requiredVersusTrophies"`
	RequiredTownhallLevel  int  `json:"requiredTownhallLevel"`
	IsWarLogPublic         bool `json:"isWarLogPublic"`
}
