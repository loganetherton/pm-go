package logging

import (
	"github.com/loganetherton/pm-go/config"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

var textFormatter log.TextFormatter
var jsonFormatter log.JSONFormatter
var jsonPrettyFormatter log.JSONFormatter
var errFormatter log.JSONFormatter

var textLog *log.Logger
var jsonLog *log.Logger
var jsonPrettyLog *log.Logger
var errLog *log.Logger

func Init() {
	SetLevel()
	CreateFormatters()
	CreateLoggers()
}

func SetLevel() {
	logLevel, levelErr := log.ParseLevel(config.LogLevel)
	if levelErr != nil {
		panic(levelErr)
	}
	log.SetLevel(logLevel)
	log.Infof("Setting log level to %s", logLevel.String())
}

func CreateLoggers() {
	textLog = log.New()
	textLog.SetFormatter(&textFormatter)

	jsonLog = log.New()
	jsonLog.SetFormatter(&jsonFormatter)

	jsonPrettyLog = log.New()
	jsonPrettyLog.SetFormatter(&jsonPrettyFormatter)

	errLog = log.New()
	errLog.SetFormatter(&errFormatter)
}

func CreateFormatters() {
	textFormatter = log.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  config.DateTimeTzFormat,
		QuoteEmptyFields: true,
	}

	jsonFormatter = log.JSONFormatter{
		TimestampFormat: config.DateTimeTzFormat,
	}

	jsonPrettyFormatter = log.JSONFormatter{
		TimestampFormat: config.DateTimeTzFormat,
		PrettyPrint:     true,
	}

	errFormatter = log.JSONFormatter{
		TimestampFormat: config.DateTimeTzFormat,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			s := strings.Split(frame.Function, ".")
			funcName := s[len(s)-1]
			relPath, _ := filepath.Rel(config.BasePath, frame.File)
			return funcName, relPath
		},
		PrettyPrint: true,
	}
}
