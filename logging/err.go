package logging

import "runtime/debug"

// LogError will be used to log any unhandled errors
func LogError(err error, msg string) {
	errLog.WithError(err).Error(msg)
	textLog.Error(string(debug.Stack()))
}
