package pool

type Clan struct {
	ID                     int
	Tag                    string `json:"tag"`
	Name                   string `json:"name"`
	Type                   string `json:"type"`
	Description            string `json:"description"`
	ClanLevel              int    `json:"clanLevel"`
	ClanPoints             int    `json:"clanPoints"`
	ClanVersusPoints       int    `json:"clanVersusPoints"`
	RequiredTrophies       int    `json:"requiredTrophies"`
	WarWinStreak           int    `json:"warWinStreak"`
	WarWins                int    `json:"warWins"`
	WarTies                int    `json:"warTies"`
	WarLosses              int    `json:"warLosses"`
	IsWarLogPublic         bool   `json:"isWarLogPublic"`
	Members                int    `json:"members"`
	RequiredVersusTrophies int    `json:"requiredVersusTrophies"`
	RequiredTownhallLevel  int    `json:"requiredTownhallLevel"`
}
