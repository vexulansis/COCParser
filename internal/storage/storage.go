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

func NewStorage() *Storage {
	s := new(Storage)
	_, err := toml.DecodeFile("db.toml", &s.Config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	DBURI := s.URI()
	DB, err := sql.Open("postgres", DBURI)
	if err != nil {
		log.Fatal()
	}
	s.DB = DB
	return s
}

func (s *Storage) URI() string {
	var DBURI string
	DBURI = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Config.Host,
		s.Config.Port,
		s.Config.User,
		s.Config.Password,
		s.Config.Name,
	)
	return DBURI
}
func (s *Storage) PrintInfo() {
	fmt.Printf("StorageInfo:\n")
	fmt.Printf("ServerPort: %d\n", s.Config.ServerPort)
	fmt.Printf("Host: %s\n", s.Config.Host)
	fmt.Printf("Port: %d\n", s.Config.Port)
	fmt.Printf("User: %s\n", s.Config.User)
	fmt.Printf("Password: %s\n", s.Config.Password)
	fmt.Printf("Name: %s\n", s.Config.Name)
}
