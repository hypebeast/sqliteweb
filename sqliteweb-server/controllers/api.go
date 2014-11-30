package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/zenazn/goji/web"
	"gopkg.in/unrolled/render.v1"

	"github.com/hypebeast/sqliteweb/sqliteweb-server/lib/client"
	"github.com/hypebeast/sqliteweb/sqliteweb-server/lib/utils"
)

var r *render.Render
var dbClient *client.DbClient

func init() {
	r = render.New(render.Options{
		IndentJSON: true,
	})
}

// Init initializes the API controller with a DB client.
func Init(client *client.DbClient) {
	dbClient = client
}

func Info(w http.ResponseWriter, req *http.Request) {
	info, err := dbClient.Info()
	if err != nil {
		renderError(w, http.StatusInternalServerError, err)
	}

	filePath, err := filepath.Abs(dbClient.DbFile)
	if err != nil {
		filePath = ""
	}

	dbName := filepath.Base(dbClient.DbFile)
	size, _ := utils.FileSize(filePath)

	result := map[string]interface{}{
		"number_of_tables":  info.Rows[0][0],
		"number_of_indexes": info.Rows[0][1],
		"filename":          dbName,
		"fullname":          filePath,
		"size":              size,
	}
	r.JSON(w, http.StatusOK, result)
}

func Tables(w http.ResponseWriter, req *http.Request) {
	tables, err := dbClient.Tables()
	if err != nil {
		renderError(w, http.StatusInternalServerError, err)
	}

	result := map[string]interface{}{
		"tables": tables,
	}
	r.JSON(w, http.StatusOK, result)
}

func Table(c web.C, w http.ResponseWriter, req *http.Request) {
	result, err := dbClient.Table(c.URLParams["name"])
	if err != nil {
		renderError(w, http.StatusInternalServerError, err)
	}

	r.JSON(w, http.StatusOK, result.Format())
}

func TableInfo(c web.C, w http.ResponseWriter, req *http.Request) {
	result, err := dbClient.TableInfo(c.URLParams["name"])
	if err != nil {
		renderError(w, http.StatusInternalServerError, err)
	}

	data := map[string]interface{}{
		"row_count":     result.Rows[0][0],
		"indexes_count": 0,
	}

	r.JSON(w, http.StatusOK, data)
}

func TableSql(c web.C, w http.ResponseWriter, req *http.Request) {
	result, err := dbClient.TableSql(c.URLParams["name"])
	if err != nil {
		renderError(w, http.StatusInternalServerError, err)
	}

	data := map[string]interface{}{
		"sql": result[0],
	}

	r.JSON(w, http.StatusOK, data)
}

func TableIndexes(c web.C, w http.ResponseWriter, req *http.Request) {
	result, err := dbClient.TableIndexes(c.URLParams["name"])
	if err != nil {
		renderError(w, http.StatusInternalServerError, err)
	}

	r.JSON(w, http.StatusOK, result.Format())
}

func Query(c web.C, w http.ResponseWriter, req *http.Request) {
	query := strings.TrimSpace(req.FormValue("query"))

	if query == "" {
		renderError(w, http.StatusBadRequest, errors.New("Query missing"))
		return
	}

	result, err := dbClient.Query(req.FormValue("query"))
	if err != nil {
		renderError(w, http.StatusInternalServerError, err)
		return
	}

	q := req.URL.Query()
	if len(q["format"]) > 0 {
		if q["format"][0] == "csv" {
			renderCSV(w, http.StatusOK, result.CSV())
			return
		} else if q["format"][0] == "json" {
			// Format the returned JSON instead of returning in the Result format
			r.JSON(w, http.StatusOK, result.Format())
			return
		}
	}

	r.JSON(w, http.StatusOK, result)
}

func OpenDatabase(c web.C, w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

// renderError renders a JSON response with the given error message.
func renderError(w http.ResponseWriter, status int, err error) {
	result := map[string]interface{}{
		"code":    "error",
		"message": err.Error(),
	}
	r.JSON(w, status, result)
}

func renderCSV(w http.ResponseWriter, status int, data []byte) {
	head := render.Head{
		ContentType: "text/csv",
		Status:      status,
	}

	d := render.Data{
		Head: head,
	}

	r.Render(w, d, data)
}
