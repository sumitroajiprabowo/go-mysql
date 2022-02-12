package gomysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	script := "insert into CUSTOMER(id, name) VALUES (3, 'Budi')"
	_, err := db.ExecContext(context.Background(), script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New Customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	script := "select * from CUSTOMER"
	rows, err := db.QueryContext(context.Background(), script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id, "name: ", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	script := "select id, name, email, balance, rating, birth_date, married, create_at from CUSTOMER"
	rows, err := db.QueryContext(context.Background(), script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		// variabel email with null string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date, create_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &create_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("==============================")
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		if email.Valid {
			fmt.Println("email: ", email.String)
		} else {
			fmt.Println("email: ", "null")
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		fmt.Println("birthDate: ", birth_date)
		fmt.Println("married: ", married)
		fmt.Println("create_at: ", create_at)
	}
}
