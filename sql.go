package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"
)

func init() {
	loadConfigGlobal()
}

func (s *Server) initDb() {
	var err error
	format := "user=%s password=%s dbname=%s host=%s"
	dataSource := fmt.Sprintf(format, config.DB.User, config.DB.Password, config.DB.Name, config.DB.Host)
	s.DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatalln("Cannot open sql. Message:", err)
		return
	}
}
