build:
	GOOS=js GOARCH=wasm go build -o ./repl/main.wasm ./repl

serve: build
	python3 -m http.server
