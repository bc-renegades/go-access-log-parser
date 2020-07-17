package database

import "os"

type Config struct {
	port     string
	host     string
	password string
	database string
	user     string
}

func NewConfig(port, host, password, database, user string) *Config {
	return &Config{port: port, host: host, password: password, database: database, user: user}
}

func NewConfigEnv() *Config {
	return NewConfig(
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_USER"),
	)
}
