bindata:
	go-bindata -o ./assets.go static/...

bindata-debug:
	go-bindata -debug -o ./assets.go static/...

setup:
	go get ./...

dev: bindata-debug
	go build -o sqliteweb .
	@echo "You can now execute ./sqliteweb"

build: bindata
	go build -o sqliteweb .

run: bindata build
	./sqliteweb

clean:
	rm ./assets.go
	rm ./spliteweb

clean-assets:
	rm static/*.html; rm -rf static/{vendor,js,css}
