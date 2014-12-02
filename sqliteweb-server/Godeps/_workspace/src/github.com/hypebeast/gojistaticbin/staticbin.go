package gojistaticbin

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// Staticbin returns a middleware handler that serves static files in the given dir
// with the help of the asset function from go-bindata (https://github.com/jteeuwen/go-bindata)
// generated file.
func Staticbin(dir string, asset func(string) ([]byte, error), options ...Options) func(http.Handler) http.Handler {
	opts := prepareOptions(options)
	modtime := time.Now()

	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, req *http.Request) {
			if req.Method != "GET" && req.Method != "HEAD" {
				h.ServeHTTP(w, req)
				return
			}

			url := req.URL.Path

			b, err := asset(filepath.Join(dir, url))

			if err != nil {
				// Try to serve the index file
				b, err = asset(filepath.Join(dir, url, opts.IndexFile))
				if err != nil {
					h.ServeHTTP(w, req)
					return
				}
			}

			if !opts.SkipLogging {
				log.Println("[STATIC] serving " + url)
			}

			http.ServeContent(w, req, url, modtime, bytes.NewReader(b))
		}

		return http.HandlerFunc(fn)
	}
}
