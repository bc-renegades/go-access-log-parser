package main

import (
	"fmt"
	"github.com/bc-renegades/go-access-log-parser/infra/database"
	"os"

	"github.com/bc-renegades/go-access-log-parser/parser"
)

func main() {
	file, err := os.Open("./logs/api.boacompra.com_access_sample.log")
	if err != nil {
		panic(err)
	}

	logs, err := parser.Parse(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("logs", logs)

	database.Open()
}
