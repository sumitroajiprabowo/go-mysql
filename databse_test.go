package gomysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:kmzway87aa@tcp(localhost:3306)/belajar_go")
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	defer db.Close()
}
