package gojistaticbin

// Options defines the options for the middleware.
type Options struct {
	// SkipLogging will disable log messages when a static file is served
	SkipLogging bool
	// IndexFile is the file that is served as an index if it exists.
	IndexFile string
}

// prepareOptions sets the default options.
func prepareOptions(options []Options) Options {
	var opts Options

	if len(options) > 0 {
		opts = options[0]
	}

	if opts.IndexFile == "" {
		opts.IndexFile = "index.html"
	}

	return opts
}
