package dbcommodi

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"../logger"
)

// Get db connection
func GetSqliteDB(path string) (db *sql.DB, err error) {
	logger.LogWarning("tst....")
	db, err = sql.Open("sqlite3", path)
	if err != nil {	
		logger.LogWarning("Error when opening db file. Error & Path: ", err, path)
	}
	return
}