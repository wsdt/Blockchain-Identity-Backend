package db

import (
	"database/sql"
	"github.com/VID-Card-Backend/_config"
	"github.com/VID-Card-Backend/controllers/errorHandling"
	_ "github.com/go-sql-driver/mysql"
)

var DB = initializeDb()
var config = _config.GetConfigDb()

func initializeDb() sql.DB {
	db, err := sql.Open("mysql",
		config.ConnStr)
	errorHandling.LogErr(err)

	if !isDbAlive(*db) {
		panic("initializeDb: DB is not alive!")
	}
	defer db.Close() // -> db connections are long lived (don't kill them), only after ending the application
	return *db
}

func isDbAlive(db sql.DB) bool {
	return db.Ping() == nil
}

/** @param query: Simple SQL query
	@param rows: Callback function. Can look like this:

		err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name)
*/
func Query(query string, cb func(rows sql.Rows)) {
	rows, err := DB.Query(query)
	errorHandling.LogErr(err)
	defer rows.Close()
	for rows.Next() {
		cb(*rows) // call callback with rows
	}
	errorHandling.LogErr(rows.Err())
}

/**
	@param escapedQuery: e.g.: select id, name from users where id = ?
	@param args: Your params for your escaped query.
*/
func SecureQuery(escapedQuery string, cb func(rows sql.Rows), args ...interface{}) {
	stmt, err := DB.Prepare(escapedQuery)
	errorHandling.LogErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(args) // Also QueryRow() possible, but by this template method we don't need to care
	errorHandling.LogErr(err)
	defer rows.Close()

	for rows.Next() {
		cb(*rows)
	}

	errorHandling.LogErr(rows.Err())
	errorHandling.LogErr(err)
}
