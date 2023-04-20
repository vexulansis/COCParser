package storage

import (
	"fmt"

	"github.com/vexulansis/COCParser/pkg/api"
)

func (s *Storage) AddClan(clan *api.Clan) error {
	if clan != nil {
		_, err := s.DB.Exec("insert into tag_name(tag,name) values ($1,$2)", clan.Tag, clan.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Storage) AddTag(tag string) error {
	_, err := s.DB.Exec("insert into tag_name(tag) values ($1)", tag)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	return nil
}
