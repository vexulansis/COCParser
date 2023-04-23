package api

import (
	db "github.com/vexulansis/COCParser/internal/storage"
)

var updquery = `update clans set
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
requiredtownhalllevel=($15) where tag=($16)`

func UpdateClanInfo(clan *Clan, storage db.Storage) error {
	_, err := storage.DB.Exec(updquery,
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
		clan.Tag,
	)
	if err != nil {
		return err
	}
	return nil
}
