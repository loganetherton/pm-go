package config

import "os"

var DbUser = os.Getenv("DB_USER")
var DbPass = os.Getenv("DB_PASS")
var DbHost = os.Getenv("DB_HOST")
var DbName = os.Getenv("DB_NAME")
