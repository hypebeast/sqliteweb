# sqliteweb

[![views](https://sourcegraph.com/api/repos/github.com/hypebeast/sqliteweb/.counters/views.svg)](https://sourcegraph.com/github.com/hypebeast/sqliteweb)
[![views 24h](https://sourcegraph.com/api/repos/github.com/hypebeast/sqliteweb/.counters/views-24h.svg)](https://sourcegraph.com/github.com/hypebeast/sqliteweb?no-count=1)

`sqliteweb` is a web-based SQLite database browser. It's written in *Go*.

It's inspired and influenced by [pgweb](https://github.com/sosedoff/pgweb) and [sqlite-browser](https://github.com/coleifer/sqlite-browser).

![](http://sebastianruml.name/images/projects/sqliteweb/sqliteweb-structure.png)

## Overview

`sqliteweb` is a web-based SQLite database browser. The goal is to provide a simple and intuitive browser for SQLite databases. It provides basic functions to browse SQLite databases and tables.

## Features

  * Browse existing SQLite databases
  * Browse table structure and indexes
  * Browse table data
  * Run custom SQL queries
  * Export query results to *CSV* and *JSON*
  * Single executable (just download the executable and run it)
  * Cross-platform

## Installation

Use the pre-build relase images from [Github Releases](https://github.com/hypebeast/sqliteweb/releases).

Currently pre-build releases are available for the following platforms:

  * Mac OSX 64bit

### From source

If there are no pre-build images for your platform, you can build sqliteweb from sources.

Requirements:

  * Go
  * Node/NPM
  * `$GOPATH` must be set

Make sure you have installed *Go*. Go +1.3 is required. You can install Go with *homebrew*:

```
$ brew install go
```

Get *sqliteweb* from Github:

```
$ git clone https://github.com/hypebeast/sqliteweb.git
$ cd sqliteweb
```

Install required packages to build the frontend:

```
$ npm install
$ bower install
```

Build the frontend:

```
$ gulp dist
```

Build the server:

```
$ cd sqliteweb-server
$ make setup
$ make build
```

Now, run *sqliteweb*:

```
$ ./sqliteweb --db ../data/test.db
```

## Usage

Start *sqliteweb* with the following command:

```
$ sqliteweb --db path/to/database/name.db
```

You can also provide a username and password to enable HTTP basic auth:

```
$ sqliteweb --db=path/to/database/name.db --auth-user=username --auth-pass=validpass
```

## CLI Options

```
$ sqliteweb -h
Usage:
  sqliteweb [OPTIONS]

Application Options:
  -v, --version    Print version
  -d, --debug      Enable debugging mode (false)
      --db=        SQLite database file
      --bind=      HTTP server host (localhost)
      --listen=    HTTP server listen port (8000)
      --auth-user= HTTP basic auth user
      --auth-pass= HTTP basic auth password
  -s, --skip-open  Skip open sqliteweb in browser on start

Help Options:
  -h, --help       Show this help message
```

## Development

The project consists of two sub-projects:

  * **sqliteweb-server**: The server part
  * **sqliteweb-web**: The frontend part

### sqliteweb-web

This is the frontend part of `sqliteweb`. The asset compilation is handled by `Gulp`.

To install all requirements, run the following commands:

```
$ cd sqliteweb
$ npm install
$ bower install
```

There is a built-in webserver for the frontend development. Run the following command to start it:

```
$ gulp
```

The compile, watch and copy the frontend files to `sqliteweb-server`:

```
$ gulp dev
```

To generate the dist files and copy them over to `sqliteweb-server`:

```
$ gulp dist
```

### sqliteweb-server

The `sqliteweb-server` provides the API for the frontend.

To build the server:

```
$ cd sqliteweb/sqliteweb-server
$ make setup
$ make dev
$ ./sqliteweb
```

## Screenshots

### Table Structure

![](http://sebastianruml.name/images/projects/sqliteweb/sqliteweb-structure.png)

### Table Content

![](http://sebastianruml.name/images/projects/sqliteweb/sqliteweb-content.png)

### SQL Query

![](http://sebastianruml.name/images/projects/sqliteweb/sqliteweb-query.png)

## Contributions

  * Fork repository
  * Create feature- or bugfix-branch
  * Create pull request
  * Use Github Issues

## Contact

  * Sebastian Ruml, <sebastian@sebastianruml.name>
  * Twitter: https://twitter.com/dar4_schneider

## License

```
The MIT License (MIT)

Copyright (c) 2014 Sebastian Ruml

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
