//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/kkrypt0nn/overscry/models"
	"github.com/kkrypt0nn/overscry/settings"
)

func parseYAML(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return js.ValueOf("Missing YAML input")
	}

	yamlStr := args[0].String()

	s, err := settings.LoadSettingsFromString(yamlStr)
	if err != nil {
		return js.ValueOf("Error parsing YAML: " + err.Error())
	}

	res := fmt.Sprintf(models.JSON_OUTPUT, s.Node.ToOQL())
	return js.ValueOf(res)
}

func main() {
	js.Global().Set("parseYAMLToOQL", js.FuncOf(parseYAML))
	select {}
}
