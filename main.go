package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bc-renegades/go-access-log-parser/infra/database"
	"github.com/bc-renegades/go-access-log-parser/parser"
	"github.com/google/uuid"
)

func main() {
	file, err := os.Open(os.Getenv("FILE_PATH"))
	if err != nil {
		panic(err)
	}

	logs, err := parser.Parse(file)
	if err != nil {
		fmt.Println(err)
	}

	c := database.NewConfigEnv()

	db := database.NewMySQLConnection(c)
	for _, log := range logs {
		db.Exec("INSERT INTO logs VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
			uuid.New(),
			log.IP,
			log.Port,
			log.Date,
			log.Resource,
			log.Method,
			log.Protocol,
			log.StatusCode,
			time.Now(),
		)
	}

	db.Close()
}
