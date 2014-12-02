/*
Package gojistatic provides a middleware for Goji that serves static files from
a go-bindata generated asset file.

Installation

To install the library, use the following command:

    go get github.com/hypebeast/gojistaticbin

Usage

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
*/
package gojistaticbin
