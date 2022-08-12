package logging

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

// LogError is the default error logger, especially for unhandled errors
func LogError(err error, msg string) {
	errLog.WithError(err).Error(msg)
	stack := debug.Stack()
	stringStack := string(stack)
	textLog.Error(stringStack)
}

// LogCallers logs out the preceding call stack
func LogCallers() {
	var callerInfo string
	pcs := make([]uintptr, 20)
	callers := runtime.Callers(0, pcs)
	pcs = pcs[:callers]

	m := make(map[string]int, len(pcs))
	frames := runtime.CallersFrames(pcs)

	for {
		frame, more := frames.Next()
		// No need to log out runtime.Callers
		if !more {
			break
		}
		fnName := frame.Function
		if strings.HasPrefix(fnName, "runtime.") {
			continue
		}
		frame.Func.Entry()
		if fnName != "" {
			//if callerInfo == "" {
			//	callerInfo = "\n"
			//}
			m[fnName] = frame.Line
			callerInfo += fmt.Sprintf("\t%v:%d\n", fnName, frame.Line)
		}
	}

	textLog.Info("Logging callers")
	plainTextLog.Info(callerInfo)
}
