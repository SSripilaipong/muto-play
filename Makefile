build:
	GOOS=js GOARCH=wasm go build -o main.wasm

serve: build
	python3 -m http.server
