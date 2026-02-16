package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL() *sql.DB {

	dsn := "root:root@tcp(mysql:3306)/authdb?parseTime=true"

	var db *sql.DB
	var err error

	for i := 0; i < 15; i++ {

		db, err = sql.Open("mysql", dsn)
		if err == nil && db.Ping() == nil {
			log.Println("Connected to MySQL")
			return db
		}

		log.Println("Waiting for MySQL...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("MySQL never became ready")
	return nil
}
