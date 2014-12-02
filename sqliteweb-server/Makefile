help:
	@echo "  setup           setup the development environment and installs all dependencies"
	@echo "  bindata         create a release asset file with go-bindata"
	@echo "  bindata-debug   create a debug asset file with go-bindata"
	@echo "  dev             build with debug asset file"
	@echo "  build           build with release asset file"
	@echo "  run             build and run sqliteweb"
	@echo "  release         build releases with gox"
	@echo "  clean           clean up"
	@echo "  clean-assets    clean up static files"

bindata:
	go-bindata -o ./assets.go static/...

bindata-debug:
	go-bindata -debug -o ./assets.go static/...

setup:
	go get github.com/tools/godep
	godep get github.com/mitchellh/gox
	godep get github.com/jteeuwen/go-bindata/...
	godep restore

dev: bindata-debug
	go build -o sqliteweb .
	@echo "You can now execute ./sqliteweb"

build: bindata
	go build -o sqliteweb .
	@echo "You can now execute ./sqliteweb"

run: bindata build
	./sqliteweb

release: bindata
	gox -osarch="darwin/amd64" -output="./bin/sqliteweb_{{.OS}}_{{.Arch}}"

clean: clean-assets
	rm ./assets.go \
	rm ./spliteweb \
	rm -rf ./bin

clean-assets:
	rm static/*.html \
	rm -rf static/vendor \
	rm -rf static/js \
	rm -rf static/css \
	rm -rf static/fonts
