package client

const (
	SQLITE_INFO = `SELECT * FROM
        (SELECT COUNT (*) AS count FROM sqlite_master WHERE type='table') AS count_tables,
        (SELECT COUNT (*) AS count FROM sqlite_master WHERE type='index') AS count_indexes;`
	SQLITE_TABLES        = "SELECT name FROM sqlite_master WHERE type='table';"
	SQLITE_TABLE_SCHEMA  = "PRAGMA table_info(%s);"
	SQLITE_TABLE_INFO    = "SELECT COUNT(*) FROM %s;"
	SQLITE_TABLE_SQL     = "SELECT sql FROM sqlite_master WHERE type='table' AND name='%s'"
	SQLITE_TABLE_INDEXES = "SELECT * FROM sqlite_master WHERE type='index' AND tbl_name='%s'"
)
