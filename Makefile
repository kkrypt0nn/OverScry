.PHONY: cli cli-docker wasm website install

cli:
	go build -o dist/overscry .

cli-docker:
	go build -o dist/overscry -ldflags '-s -w' .

wasm:
	GOOS=js GOARCH=wasm go build -o dist/overscry.wasm -ldflags '-s -w' ./wasm
	cp dist/overscry.wasm web/online
	cp "$(shell go env GOROOT)/lib/wasm/wasm_exec.js" web/online

website:
	$(MAKE) wasm
	-mkdir www
	cd web/book && mdbook build --dest-dir ../../www
	-mkdir www/online
	cp web/online/* www/online
	cd www && npx sscli --no-clean --base https://overscry.krypton.ninja
	cd www/online && npx @tailwindcss/cli -i input.css -o output.css && rm input.css

install:
	go install .
