package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Port string
var Domain string
var MySQLDB string
var MySQLUser string
var MySQLPassword string
var DBHost string

func LoadEnv() {
	
	godotenv.Load()
  
	Port = os.Getenv("PORT")
	MySQLDB = os.Getenv("MYSQL_DATABASE")
	MySQLUser = os.Getenv("MYSQL_USER")
	MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	DBHost = os.Getenv("DB_HOST")
	Domain = os.Getenv("DOMAIN")
}