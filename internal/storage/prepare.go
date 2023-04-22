package storage

func (s *Storage) Prepare(n int) error {
	for i := 0; i < n; i++ {
		_, err := s.DB.Exec("insert into clans default values")
		if err != nil {
			return err
		}
	}
	return nil
}
