package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type addConnectionRequest struct {
	address  string
	username string
	password string
}

type connectionModelDB struct {
	id       int32
	address  string
	username string
	password string
}

func dbConnect(addr, user, pass string) (*sql.DB, error) {
	mysqlCfg := mysql.Config{
		User:   user,
		Passwd: pass,
		Addr:   addr,
		DBName: "ssh",
		Net:    "tcp",
	}

	db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error open db connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error during ping to db: %v", err)
	}

	return db, nil
}

func (s service) addConnectionDB(conn addConnectionRequest) error {
	_, err := s.db.Exec("INSERT INTO connections (address, username, password) VALUES (?, ?, ?)",
		conn.address, conn.username, conn.password)
	if err != nil {
		return fmt.Errorf("addConnection error: %v", err)
	}

	return nil
}

func (s service) getConnectionByIdDB(id string) (connectionModelDB, error) {
	var conn connectionModelDB

	row := s.db.QueryRow("SELECT * from connections where id = ?", id)
	err := row.Scan(&conn.id, &conn.address, &conn.username, &conn.password)
	if err != nil {
		return connectionModelDB{}, fmt.Errorf("get connection by id error: %v", err)
	}

	return conn, nil
}

func (s service) deleteConnectionDB(id string) (connectionModelDB, error) {
	conn, err := s.getConnectionByIdDB(id)
	if err != nil {
		return connectionModelDB{}, err
	}

	_, err = s.db.Query("DELETE FROM connections WHERE id = ?", id)
	if err != nil {
		return connectionModelDB{}, fmt.Errorf("cannot delete connection with id [%s]: %v", id, err)
	}

	return conn, nil
}

func (s service) listConnectionsDB() ([]connectionModelDB, error) {
	var conns []connectionModelDB

	rows, err := s.db.Query("SELECT * FROM connections")
	if err != nil {
		return nil, fmt.Errorf("query select error: %v", err)
	}
	for rows.Next() {
		var conn connectionModelDB
		err = rows.Scan(&conn.id, &conn.address, &conn.username, &conn.password)
		if err != nil {
			return nil, fmt.Errorf("get connection by host error: %v", err)
		}
		conns = append(conns, conn)
	}

	return conns, nil
}
