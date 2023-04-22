package storage

import (
	"fmt"

	"github.com/vexulansis/COCParser/pkg/api"
)

func (s *Storage) GenerateTags(n int) error {
	for i := 0; i < n; i++ {
		_, err := s.DB.Exec("update clans set tag=($1) where id=($2)", TagFromId(i), i)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Storage) AddClanInfo(n int) error {
	for i := 500; i < n; i++ {
		clan, _ := api.GetClanByTag(TagFromId(i))
		if clan != nil {
			querystr := `update clans set
		name=($1),
		type=($2),
		description=($3),
		clanlevel=($4),
		clanpoints=($5),
		clanversuspoints=($6),
		requiredtrophies=($7),
		warwinstreak=($8),
		warwins=($9),
		warties=($10),
		warlosses=($11),
		iswarlogpublic=($12),
		members=($13),
		requiredversustrophies=($14),
		requiredtownhalllevel=($15) where id=($16)`

			_, err := s.DB.Exec(querystr,
				clan.Name,
				clan.Type,
				clan.Description,
				clan.ClanLevel,
				clan.ClanPoints,
				clan.ClanVersusPoints,
				clan.RequiredTrophies,
				clan.WarWinStreak,
				clan.WarWins,
				clan.WarTies,
				clan.WarLosses,
				clan.IsWarLogPublic,
				clan.Members,
				clan.RequiredVersusTrophies,
				clan.RequiredTownhallLevel,
				i,
			)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return err
			}
		}
	}
	return nil
}
