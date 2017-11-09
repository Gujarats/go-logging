package logging

import (
	"os"
)

// Direct call to log the message
// So you don't have to set the initial variable
func Debug(prefix, message string) {

	logBackend := NewLogBackend(os.Stderr, prefix, 0)
	SetBackend(logBackend)
	format := MustStringFormatter(
		`%{color}%{time:15:04:05.000} ▶ %{color:reset} %{message}`,
	)
	SetFormatter(format)
	log := MustGetLogger("test")
	log.Debug(message)
}
