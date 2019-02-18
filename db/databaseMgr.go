package db

import (
	"database/sql"
	"github.com/VID-Card-Backend/_config"
	"github.com/VID-Card-Backend/controllers/errorHandling"
	_ "github.com/go-sql-driver/mysql"
)

var DB sql.DB = initializeDb()
var config = _config.GetConfigDb()

func initializeDb() sql.DB {
	db, err := sql.Open("mysql",
		config.ConnStr)
	errorHandling.LogErr(err)

	if !isDbAlive(*db) {
		panic("initializeDb: DB is not alive!")
	}
	return *db
	// defer db.Close() -> db connections are long lived (don't kill them), only after ending the application
}

func isDbAlive(db sql.DB) bool {
	return db.Ping() == nil
}
