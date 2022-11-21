package Config

import (
	"database/sql"
	"os"
)

func InitDB() *sql.DB {
	db, err := sql.Open(os.Getenv("DBDRIVER"), os.Getenv("DBURL"))
	if err != nil {
		panic(err)
	}

	return db
}
