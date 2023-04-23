package storage

import "github.com/vexulansis/COCParser/internal/algo"

func (s *Storage) Prepare(n int) error {
	for i := 0; i < n; i++ {
		_, err := s.DB.Exec("insert into clans default values")
		if err != nil {
			return err
		}
	}
	s.GenerateTags(0, n)
	return nil
}
func (s *Storage) GenerateTags(start int, end int) error {
	for i := start; i < end; i++ {
		_, err := s.DB.Exec("update clans set tag=($1) where id=($2)", algo.TagFromId(i), i)
		if err != nil {
			return err
		}
	}
	return nil
}
