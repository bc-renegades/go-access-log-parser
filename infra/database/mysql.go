package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(c *Config) *sql.DB {
	fmt.Printf("%s:%s@tcp(%s:%s)/%s", c.user, c.password, c.host, c.port, c.database)
	db, err := sql.Open(c.host, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.user, c.password, c.host, c.port, c.database))
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database CONNECTED")

	return db
}
