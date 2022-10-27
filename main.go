package main

import (
	"database/sql"
	"log"
	"net/http"
)

type service struct {
	db  *sql.DB
	cfg config
}

func createService() (service, error) {
	cfg, err := createConfig()
	if err != nil {
		return service{}, err
	}

	db, err := dbConnect(cfg.dbAddr, cfg.dbUser, cfg.dbPass)
	if err != nil {
		return service{}, err
	}

	return service{db: db, cfg: cfg}, nil
}

func (s service) run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/connection", s.Connection)
	mux.HandleFunc("/connections", s.Connections)

	if err := http.ListenAndServe(s.cfg.httpAddr, mux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	s, err := createService()
	if err != nil {
		log.Fatal(err)
	}

	s.run()
}
