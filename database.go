package gomysql

import (
	"database/sql"
	"fmt"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:kmzway87aa@tcp(localhost:3306)/learn_golang?parseTime=true")
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
