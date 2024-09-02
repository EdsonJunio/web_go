package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {

	dsn := "root:270519@tcp(localhost:3306)/loja"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
