package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Open() {
	db, err := sql.Open("mysql", "root:dev@tcp(mysql:3306)/renegades")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("BANCO CONNECTED")
}
