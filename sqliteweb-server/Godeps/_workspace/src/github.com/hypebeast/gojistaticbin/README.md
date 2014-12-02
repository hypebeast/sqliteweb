# goji-staticbin

goji-staticbin is a middleware for [Goji](https://github.com/zenazn/goji) that serves static files from a [go-bindata](https://github.com/jteeuwen/go-bindata) generated asset file.

[![GoDoc](https://godoc.org/github.com/hypebeast/gojistaticbin?status.svg)](https://godoc.org/github.com/hypebeast/gojistaticbin)


## Installation

To install the library, use the following command:

```
go get github.com/hypebeast/gojistaticbin
```


## Usage

```go
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

```

See the `examples` folder for an example application that uses goji-staticbin.

