package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("sqlite3", "./configs/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Database connect")
}
