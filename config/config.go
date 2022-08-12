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
var appEnv = os.Getenv("APP_ENV")
var IsDev = appEnv == "dev" || appEnv == "development"
var IsTest = appEnv == "test"
var IsProd = appEnv == "prod" || appEnv == "production"

// LogLevel is used to set the level of logging
var LogLevel = os.Getenv("LOG_LEVEL")
