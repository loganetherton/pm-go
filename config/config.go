package config

import (
	"os"
	"path/filepath"
)

var cwd, _ = os.Getwd()
var BasePath = filepath.Join(cwd, "..")

/*
Logging configuration
*/
var LogLevel = os.Getenv("LOG_LEVEL")

/*
Common formats
*/
var DateFormat = "2006-01-02"
var DateTimeFormat = "2006-01-02 15:04:05"
var DateTimeTzFormat = "2006-01-02 15:04:05 -0700"
