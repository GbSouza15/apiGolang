package api

import "database/sql"

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}
