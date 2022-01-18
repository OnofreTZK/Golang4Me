package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {

	db, err := sql.Open("mysql", "root:MetamaxT08@tcp(localhost)/")
	if err != nil {
		panic(err)
	}
	// Right before the main function finished the database connection will be closed
	defer db.Close()

	exec(db, "create database if not exists golang")

}
