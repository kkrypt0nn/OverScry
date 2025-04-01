package main

import (
	"fmt"
	"runtime"

	"github.com/kkrypt0nn/overscry/cmd"
	"github.com/kkrypt0nn/overscry/core"
)

func main() {
	core.Logger.Println(fmt.Sprintf("%s v%s ${fg:gray}(built for %s on %s)${effect:reset}\n", core.Name, core.Version, runtime.GOOS, runtime.GOARCH))
	cmd.Execute()
}
