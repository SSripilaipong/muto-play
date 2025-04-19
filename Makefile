build: build-repl build-ide

build-ide:
	GOOS=js GOARCH=wasm go build -o ./ide/ide.wasm ./ide

build-repl:
	GOOS=js GOARCH=wasm go build -o ./repl/repl.wasm ./repl

serve: build
	python3 -m http.server

update-muto: go-get-muto build

go-get-muto:
	go get github.com/SSripilaipong/muto@$(shell curl -s https://api.github.com/repos/SSripilaipong/muto/commits/main | jq -r .sha)
