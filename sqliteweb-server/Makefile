help:
	@echo "  env         create a development environment using virtualenv"
	@echo "  deps        install dependencies"
	@echo "  clean       remove unwanted stuff"
	@echo "  lint        check style with flake8"
	@echo "  production  Clean and do a release"
	@echo "  release     package and upload a release"
	@echo "  sdist       package"

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
