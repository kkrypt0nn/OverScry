.PHONY: cli cli-docker wasm install

cli:
	go build -o dist/overscry .

cli-docker:
	go build -o dist/overscry -ldflags '-s -w' .

documentation:
	mdbook build --dest-dir web/book

wasm:
	GOOS=js GOARCH=wasm go build -o dist/overscry.wasm ./wasm
	cp dist/overscry.wasm web/wasm
	cp "$(shell go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm

install:
	go install .
