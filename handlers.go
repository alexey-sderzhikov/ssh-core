package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s service) Connection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.GetConnection(w, r)
	case "POST":
		s.AddConnection(w, r)
	case "DELETE":
		s.DeleteConnection(w, r)
	}
}

func (s service) Connections(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.ListConnections(w, r)
		// case "POST":
		// 	AddConnection(w, r)
		// case "DELETE":
		// 	DeleteConnection(w, r)
	}
}

func (s service) GetConnection(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error during parse body: %v", err)
	}
	id := r.URL.Query().Get("id")

	connDB, err := s.getConnectionByIdDB(id)
	if err != nil {
		fmt.Println(err)
	}

	connHTTP := mapConnectionDBtoHTTP(connDB)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connHTTP)
}

func (s service) AddConnection(w http.ResponseWriter, r *http.Request) {
	var conn addConnectionRequest

	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error during parse body: %v", err)
	}
	conn.address = r.Form.Get("address")
	conn.username = r.Form.Get("username")
	conn.password = r.Form.Get("password")

	s.addConnectionDB(conn)
}

func (s service) DeleteConnection(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error during parse body: %v", err)
	}
	id := r.URL.Query().Get("id")

	connDB, err := s.deleteConnectionDB(id)
	if err != nil {
		fmt.Println(err)
	}

	connHTTP := mapConnectionDBtoHTTP(connDB)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connHTTP)
}

func (s service) ListConnections(w http.ResponseWriter, r *http.Request) {
	consDB, err := s.listConnectionsDB()
	if err != nil {
		fmt.Println(err)
	}

	var consHTTP connectionListHTTP

	consHTTP.Conns = mapConnectionsListDBtoHTTP(consDB)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consHTTP)
}

type connectionHTTP struct {
	ID       int32  `json:"id"`
	Address  string `json:"address"`
	Username string `json:"username"`
}

type connectionListHTTP struct {
	Conns []connectionHTTP `json:"connectioins"`
}
