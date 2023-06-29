package pkg

import (
	"database/sql"
	"fmt"
)

func StartDB() *sql.DB {
	db, err := InitPostgres()

	if err != nil {
		fmt.Println(err)
		db.Close()
		return db
	}
	return db
}
