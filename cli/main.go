package cli

import (
	"fmt"
	"runtime"

	"github.com/kkrypt0nn/overscry/cli/cmd"
	"github.com/kkrypt0nn/overscry/core"
)

func Main() {
	core.Logger.Println(fmt.Sprintf("%s v%s ${fg:gray}(built for %s on %s)${effect:reset}\n", core.Name, core.Version, runtime.GOOS, runtime.GOARCH))
	cmd.Execute()
}
