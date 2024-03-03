package main

import (
	"github.com/loganetherton/pm-go/db"
	"github.com/loganetherton/pm-go/logging"
	"github.com/loganetherton/pm-go/utils"
	"github.com/loganetherton/pm-go/web"
)

func main() {
	defer utils.Recover()
	go logging.Init()
	db.Connect()
	web.Init()
}
