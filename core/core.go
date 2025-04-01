package core

import (
	"github.com/kkrypt0nn/tangra"
)

type CustomLogger struct {
	*tangra.Logger
}

var Logger *CustomLogger

func (l *CustomLogger) Success(message string) {
	message = "${datetime} ${fg:green}SUCCESS${reset}: " + message
	l.Println(message)
}

func init() {
	Logger = &CustomLogger{
		tangra.NewLogger(),
	}
}
