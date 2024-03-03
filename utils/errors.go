package utils

import (
	"github.com/loganetherton/pm-go/logging"
	"github.com/loganetherton/pm-go/types"
)

func Recover(messages ...string) {
	var message string
	if messages != nil {
		for _, s := range messages {
			if message != "" {
				message += ";"
			}
			message += s
		}
	}
	if e := recover(); e != nil {
		if !Implements(e, types.ErrorInterface) {
			panic("Recover caught something other than an error (somehow)")
		}
		logging.LogError(e.(error), message)
	}
}
