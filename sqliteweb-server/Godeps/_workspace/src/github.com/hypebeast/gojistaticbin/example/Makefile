bindata:
	go-bindata -o ./assets.go static/...

bindatadebug:
	go-bindata -debug -o ./assets.go static/...

setup:
	go get ./...

build:
	go build -o gojistaticdata main.go assets.go

run: bindata build
	./gojistaticdata
