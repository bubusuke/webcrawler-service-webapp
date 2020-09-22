package db

import (
	"fmt"
	"log"
	"os"
)

// dbParam represents Database connecting parameters.
type dbParam struct {
	host     string
	port     string
	dbname   string
	user     string
	password string
}

// param manages dbParam with a singleton.
var param dbParam

// init reads env and create param.
func init() {
	param = dbParam{
		host:     ifBlank(os.Getenv("DB_HOST"), "localhost"),
		port:     ifBlank(os.Getenv("DB_PORT"), "5432"),
		dbname:   ifBlank(os.Getenv("DB_DATABASE_NAME"), "postgres"),
		user:     ifBlank(os.Getenv("DB_USER"), "postgres"),
		password: ifBlank(os.Getenv("DB_PASSWORD"), "pass"),
	}

	log.Println("DB information:", GetDbInfo())
}

// GetDbInfo returns database connection string.
func GetDbInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		param.host, param.port, param.user, param.password, param.dbname)
}

func ifBlank(val, defaultVal string) string {
	if val == "" {
		return defaultVal
	}
	return val
}
