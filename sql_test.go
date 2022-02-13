package gomysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin'; #"
	password := "salah"

	q := "select * from users where username = '" + username + "' and password = '" + password + "' limit 1"
	fmt.Println(q)
	rows, err := db.QueryContext(context.Background(), q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin"
	password := "admin"

	q := "select * from users where username = ? and password = ? limit 1"
	fmt.Println(q)
	rows, err := db.QueryContext(context.Background(), q, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "danu"
	password := "danu"

	q := "insert into users(username, password) VALUES (?, ?)"
	_, err := db.ExecContext(context.Background(), q, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New Users")
}

func TestAutoincrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "danu@anakdesa.id"
	comment := "Golang"

	q := "insert into comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(context.Background(), q, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New Comments with Id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	q := "insert into comments(email, comment) VALUES (?, ?)"
	stmt, err := db.PrepareContext(context.Background(), q)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 20; i++ {
		email := "danu" + strconv.Itoa(i) + "@anakdesa.id"
		comment := "Komentar ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(context.Background(), email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comments with Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	q := "insert into comments(email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {
		email := "danu" + strconv.Itoa(i) + "@anakdesa.id"
		comment := "Komentar ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(context.Background(), q, email, comment)
		if err != nil {
			tx.Rollback()
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println("Comments with Id", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}
