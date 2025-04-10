.PHONY: cli cli-docker wasm website install

cli:
	go build -o dist/overscry .

cli-docker:
	go build -o dist/overscry -ldflags '-s -w' .

wasm:
	GOOS=js GOARCH=wasm go build -o dist/overscry.wasm ./wasm
	cp dist/overscry.wasm web/wasm
	cp "$(shell go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm

website:
	$(MAKE) wasm
	-mkdir www && mkdir www/wasm
	cp web/wasm/* www/wasm
	cd web/book && mdbook build
	cp -R web/book/book/ www

install:
	go install .
