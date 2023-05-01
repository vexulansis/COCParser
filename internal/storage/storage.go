package storage

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	Logger *logrus.Logger
	DB     *sql.DB
	Config DBConfig
}
type DBConfig struct {
	Host        string
	User        string
	Password    string
	Name        string
	APIemail    string
	APIpassword string
	ServerPort  int
	Port        int
}

func NewStorage() (*Storage, error) {
	s := new(Storage)
	// Custom logrus logger
	logger := initLogger()
	s.Logger = logger
	// Get DB info from .toml
	_, err := toml.DecodeFile("db.toml", &s.Config)
	if err != nil {
		s.Logger.Fatal(err)
	}
	// Create DB URI
	DBURI := s.URI()
	// Open DB
	DB, err := sql.Open("postgres", DBURI)
	if err != nil {
		s.Logger.Fatal(err)
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
