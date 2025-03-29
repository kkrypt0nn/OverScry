package core

import "github.com/kkrypt0nn/tangra"

var Logger *tangra.Logger

func init() {
	Logger = tangra.NewLogger()
}
