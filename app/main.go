package main

import (
	"ShelfQL/cmd"
	"database/sql"
	"log"
)

func main() {
	cmd.Execute()

	db := createHandle()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing db connection: %v", err)
		}
	}(db)

	insertData(db, "The Book Thief")
}
