package db

//здесь планируется доставать из бд все, что связано с сигнатурами и тд

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect_db() {

	connStr := "user=root password=root dbname=av_signatures sslmode=disabled"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
