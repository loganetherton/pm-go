package config

import (
	"os"
	"path/filepath"
)

/*
Application directory paths
*/
var cwd, _ = os.Getwd()
var BasePath = filepath.Join(cwd, "..")

/*
*
Application environment
*/
var AppEnv = os.Getenv("APP_ENV")
var IsDev = AppEnv == "dev" || AppEnv == "development"
var IsTest = AppEnv == "test"
var IsProd = AppEnv == "prod" || AppEnv == "production"

// LogLevel is used to set the level of logging
var LogLevel = os.Getenv("LOG_LEVEL")
