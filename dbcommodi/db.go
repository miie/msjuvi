package dbcommodi

import (
	//"errors"
	//"reflect"
	//"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"../logger"
)

/*1. flytta all till en "db" paket och skapa GetDB för specifika databaser
2. rensa upp lite och skapa en execqueryrow för transactions 
3. Dokumentera funktionerna
4. skriv tester (se auth_test.go)*/

var ErrNoRows = sql.ErrNoRows

func ExecQuery(db *sql.DB, query string, queryargs ...interface{}) (result sql.Result, err error) {
	if queryargs == nil {
		result, err = db.Exec(query)	
	} else {
		result, err = db.Exec(query, queryargs...)
	}
	if err != nil {
		logger.LogWarning("ExecQuery: Error when executing query. Query & Error: " + query + err.Error())
	}
	return
}

func ExecQueryRow(db *sql.DB, query string, queryargs ...interface{}) (row *sql.Row) {
	// Check if we should run with or without args
	if queryargs != nil {
		row = db.QueryRow(query, queryargs...)
	} else {
		row = db.QueryRow(query)
	}
	return
}

func ScanRow(row *sql.Row, scanargs ...interface{}) (err error) {
	err = row.Scan(scanargs...)
	switch {
    case err == sql.ErrNoRows:
            return ErrNoRows
    case err != nil:
            logger.LogWarning("ScanRow: error when scanning row. error: " + err.Error())
    }
    return
}

func GetTransactionAndStatement(db *sql.DB, query string) (tx *sql.Tx, stmt *sql.Stmt, err error) {
	// Create transaction for update query 
	tx, err = db.Begin()
	if err != nil {
		logger.LogWarning("GetTransactionAndStatement: Error when starting transaction. Error: " + err.Error())
		return
	}

	// Prepare update statement
	stmt, err = tx.Prepare(query)
	if err != nil {
		logger.LogWarning("Error when preparing query for transaction. Query & Error: " + query + err.Error())
	}
	return
}


func RunQuery(db *sql.DB, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = db.Query(query, args)
	if err != nil {
		logger.LogWarning("RunQuery: Error when running query. Query & Error: " + query + err.Error())
	}
	return
}

func ExecQueryWithTransaction(stmt *sql.Stmt, args ...interface{}) (result sql.Result, err error) {
	result, err = stmt.Exec(args)
	if err != nil {
		logger.LogWarning("ExecQueryWithTransaction: Error when executing statement. Error: " + err.Error())
	}
	return
}

func RunQueryWithTransaction(stmt *sql.Stmt, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = stmt.Query(args)
	if err != nil {
		logger.LogWarning("RunQueryWithTransaction: Error when executing statement. Query & Error: " + err.Error())
	}
	return	
}