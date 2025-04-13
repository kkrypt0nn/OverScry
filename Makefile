.PHONY: cli cli-docker wasm website install

cli:
	go build -o dist/overscry .

cli-docker:
	go build -o dist/overscry -ldflags '-s -w' .

wasm:
	GOOS=js GOARCH=wasm go build -o dist/overscry.wasm -ldflags '-s -w' ./wasm
	cp dist/overscry.wasm web/wasm
	cp "$(shell go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm

website:
	$(MAKE) wasm
	-mkdir www
	cd web/book && mdbook build --dest-dir ../../www
	-mkdir www/wasm
	cp web/wasm/* www/wasm
	cd www && npx sscli --no-clean --base https://overscry.krypton.ninja

install:
	go install .
