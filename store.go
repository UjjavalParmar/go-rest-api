package main

import "database/sql"

type Store interface{

}

func NewStore(db *sql.DB) Store {
	return nil
}

