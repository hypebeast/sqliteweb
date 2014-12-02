package main

import (
	"github.com/hypebeast/gojistaticbin"
	"github.com/zenazn/goji"
)

func main() {
	goji.Use(gojistaticbin.Staticbin("static", Asset, gojistaticbin.Options{
		SkipLogging: false,
		IndexFile:   "index.html",
	}))

	goji.Serve()
}
