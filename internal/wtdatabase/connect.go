package wtdatabase

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var wtdb *sql.DB

//Initialize Database Connection, save pointer of the database and return the pointer as result
func InitDb(dbType string, dbConnectionString string) (*sql.DB, error) {
	wtdb = nil
	db, err := sql.Open(dbType, dbConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	wtdb = db
	return wtdb, nil
}

//Get database connection pointer
func GetDb() (db *sql.DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("database error : %v", r)
		}
	}()

	if wtdb == (*sql.DB)(nil) {
		return nil, fmt.Errorf("database was not initialized yet")
	}

	return wtdb, nil
}

//Get database connection pointer
func ClearDb() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("database error : %v", r)
		}
	}()
	err = nil
	if wtdb != (*sql.DB)(nil) {
		err = wtdb.Close()
		wtdb = nil
		return err
	}
	wtdb = nil
	return err
}
