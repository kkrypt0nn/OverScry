//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/kkrypt0nn/overscry/core"
	"github.com/kkrypt0nn/overscry/models"
	"github.com/kkrypt0nn/overscry/settings"
)

func convertYAMLToOQL(this js.Value, args []js.Value) any {
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

func getVersion(this js.Value, args []js.Value) any {
	return js.ValueOf(core.Version)
}

func main() {
	js.Global().Set("convertYAMLToOQL", js.FuncOf(convertYAMLToOQL))
	js.Global().Set("getVersion", js.FuncOf(getVersion))
	select {}
}
