package dbcommodi

import (
	"os"
	"testing"
)

var testdbpath string = "./db_test.db"

func TestExecQuery(t *testing.T) {

	os.Remove(testdbpath)
	db, err := GetSqliteDB(testdbpath)
	if err != nil {
		t.Errorf("ExecQueryTest: error when getting sqlite db. error & path", err, testdbpath)
	}
	defer db.Close()
	defer os.Remove(testdbpath)

	query := "CREATE TABLE auth (id TEXT NOT NULL PRIMARY KEY, chall0 TEXT, chall1 TEXT)"
	_, err = ExecQuery(db, query)
	if err != nil {
		t.Errorf("ExecQueryTest: error when creating table")
	}

	query = "INSERT INTO auth(id, chall0, chall1) values(?, ?, ?)"
	_, err = ExecQuery(db, query, "0000180290bd5aa24d47f26002cbed4c4d2d63fa", "sdfafasd", "dsalfkaslkddsöakfhslajb")
	if err != nil {
		t.Errorf("ExecQueryTest: error when inserting")
	}

	query = "SELECT chall0 FROM auth WHERE id = ?"
	row := ExecQueryRow(db, query, "0000180290bd5aa24d47f26002cbed4c4d2d63fa")
	var chall0 string
	err = row.Scan(&chall0)
	if err != nil {
		t.Errorf("ExecQueryTest: error when scanning for chall0. error: ", err)
	}

	if chall0 != "sdfafasd" {
		t.Errorf("ExecQueryTest: error when getting chall0. chall0: ", chall0)
	}

	query = "SELECT chall0 FROM auth WHERE id = ?"
	row = ExecQueryRow(db, query, "0")
	err = row.Scan(&chall0)
	if err != ErrNoRows {
		t.Errorf("ExecQueryTest: error, should get ErrNoRow. Instead got error: ", err)
	}
	
}