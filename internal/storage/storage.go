package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type Storage struct {
	Config DBConfig
	DB     *sql.DB
}
type DBConfig struct {
	ServerPort int
	Host       string
	Port       int
	User       string
	Password   string
	Name       string
}

func NewStorage() (*Storage, error) {
	s := new(Storage)
	_, err := toml.DecodeFile("db.toml", &s.Config)
	if err != nil {
		log.Fatal()
	}
	DBURI := s.URI()
	DB, err := sql.Open("postgres", DBURI)
	if err != nil {
		log.Fatal()
	}
	s.DB = DB
	return s, nil
}

func (s *Storage) URI() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Config.Host,
		s.Config.Port,
		s.Config.User,
		s.Config.Password,
		s.Config.Name,
	)
}
