package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"github.com/hypebeast/gojistaticbin"
	flags "github.com/jessevdk/go-flags"
	"github.com/zenazn/goji"

	"github.com/hypebeast/sqliteweb/sqliteweb-server/controllers"
	"github.com/hypebeast/sqliteweb/sqliteweb-server/lib/client"
	"github.com/hypebeast/sqliteweb/sqliteweb-server/lib/utils"
)

const VERSION = "0.0.1"

var options struct {
	Version  bool   `short:"v" long:"version" description:"Print version"`
	Debug    bool   `short:"d" long:"debug" description:"Enable debugging mode" default:"false"`
	DbFile   string `long:"db" description:"SQLite database file"`
	HttpHost string `long:"bind" description:"HTTP server host" default:"localhost"`
	HttpPort uint   `long:"listen" description:"HTTP server listen port" default:"8000"`
	AuthUser string `long:"auth-user" description:"HTTP basic auth user"`
	AuthPass string `long:"auth-pass" description:"HTTP basic auth password"`
	SkipOpen bool   `short:"s" long:"skip-open" description:"Skip open sqliteweb in browser on start"`
}

// initOptions initialize the command line options.
func initOptions() {
	_, err := flags.Parse(&options)
	utils.ExitOnError(err, "error parsing options")

	// TODO: Add support for env variables

	if options.Version {
		printHeader()
		os.Exit(0)
	}

	// Remove all flags to make sure that Goji doesn't read them
	os.Args = os.Args[:1]

	// TODO: Configure log to support logging to a file
}

// initClient initialize the database client.
func initClient() *client.DbClient {
	if options.DbFile == "" {
		log.Println("No database specified")
		return nil
	}

	db, err := client.New(options.DbFile)
	if err != nil {
		log.Panicf("ERROR: Can not open database: %s\n", options.DbFile)
	}

	return db
}

// initServer initialize and start the web server.
func initServer(db *client.DbClient) {
	// Initialize the API controller
	controllers.Init(db)

	goji.Use(gojistaticbin.Staticbin("static", Asset, gojistaticbin.Options{
		SkipLogging: true,
		IndexFile:   "index.html",
	}))

	// api := web.New()
	// goji.Handle("/api/*", api)

	goji.Get("/api/info", controllers.Info)
	goji.Get("/api/table", controllers.Tables)
	goji.Get("/api/table/:name", controllers.Table)
	goji.Get("/api/table/:name/info", controllers.TableInfo)
	goji.Get("/api/table/:name/sql", controllers.TableSql)
	goji.Get("/api/table/:name/indexes", controllers.TableIndexes)
	goji.Get("/api/query", controllers.Query)
	goji.Post("/api/query", controllers.Query)

	address := fmt.Sprintf("%s:%d", options.HttpHost, options.HttpPort)
	flag.Set("bind", address)

	go goji.Serve()
}

// openBrowser open the page in the browser.
func openBrowser() {
	if options.SkipOpen {
		return
	}

	_, err := exec.Command("which", "open").Output()
	if err != nil {
		log.Println("Error opening sqliteweb in browser")
		return
	}

	location := fmt.Sprintf("http://%s:%d", options.HttpHost, options.HttpPort)
	_, err = exec.Command("open", location).Output()
	if err != nil {
		log.Println("Error opening sqliteweb in browser")
	}
}

// waitForExits wait until the user exists the program.
func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

// printHeader print the welcome header.
func printHeader() {
	fmt.Fprintf(os.Stdout, "sqliteweb, v%s\n", VERSION)
}

// Everything starts from here.
func main() {
	initOptions()

	printHeader()

	db := initClient()
	initServer(db)

	openBrowser()
	waitForExit()
}
