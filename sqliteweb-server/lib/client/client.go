package client

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

type DbClient struct {
	db      *sqlx.DB
	DbFile  string
	history []string
}

type Row []interface{}

type Result struct {
	Columns []string `json:"columns"`
	Rows    []Row    `json:"rows"`
}

func New(file string) (*DbClient, error) {
	db, err := sqlx.Open("sqlite3", file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &DbClient{db: db, DbFile: file}, nil
}

func (client *DbClient) Info() (*Result, error) {
	return client.query(SQLITE_INFO)
}

func (client *DbClient) Tables() ([]string, error) {
	return client.fetchRows(SQLITE_TABLES)
}

func (client *DbClient) TableInfo(table string) (*Result, error) {
	return client.query(fmt.Sprintf(SQLITE_TABLE_INFO, table))
}

// Table returns the table structure.
func (client *DbClient) Table(table string) (*Result, error) {
	return client.query(fmt.Sprintf(SQLITE_TABLE_SCHEMA, table))
}

// TableSql returns the SQL used to create the given table.
func (client *DbClient) TableSql(table string) ([]string, error) {
	return client.fetchRows(fmt.Sprintf(SQLITE_TABLE_SQL, table))
}

// TableIndexes returns the indexes for the given table.
func (client *DbClient) TableIndexes(table string) (*Result, error) {
	return client.query(fmt.Sprintf(SQLITE_TABLE_INDEXES, table))
}

func (client *DbClient) Query(query string) (*Result, error) {
	client.recordQuery(query)
	return client.query(query)
}

func (db *DbClient) query(query string, args ...interface{}) (*Result, error) {
	rows, err := db.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := &Result{Columns: columns}

	for rows.Next() {
		cols, err := rows.SliceScan()
		if err != nil {
			continue
		}

		for i, item := range cols {
			if item == nil {
				cols[i] = nil
			} else {
				t := reflect.TypeOf(item).Kind().String()

				if t == "slice" {
					cols[i] = string(item.([]byte))
				}
			}
		}

		result.Rows = append(result.Rows, cols)
	}

	return result, nil
}

// fetchRows return a string slice of all rows for the first column in the
// query result.
func (client *DbClient) fetchRows(query string) ([]string, error) {
	res, err := client.query(query)
	if err != nil {
		return nil, err
	}

	// Init empty slice; otherwise JSON marshal will encode it to "null"
	results := make([]string, 0)

	for _, row := range res.Rows {
		results = append(results, row[0].(string))
	}

	return results, nil
}

// recordQuery adds the query to the query history.
func (client *DbClient) recordQuery(query string) {
	client.history = append(client.history, query)
}

// Format returns a slice of maps. The key in the map represents the column name
// and the value is the row content.
func (res *Result) Format() []map[string]interface{} {
	var items []map[string]interface{}

	for _, row := range res.Rows {
		item := make(map[string]interface{})
		for i, c := range res.Columns {
			item[c] = row[i]
		}

		items = append(items, item)
	}

	return items
}

func (res *Result) CSV() []byte {
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	// Write the header
	writer.Write(res.Columns)

	// Write the values
	for _, row := range res.Rows {
		record := make([]string, len(row))

		for i, val := range row {
			var v string
			if val != nil {
				v = fmt.Sprintf("%v", val)
			} else {
				v = ""
			}
			record[i] = v
		}
		writer.Write(record)
	}

	return buf.Bytes()
}
