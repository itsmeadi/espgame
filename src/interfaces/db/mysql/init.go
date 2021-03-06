package mysql

import (
	"github.com/ESPGame/src/interfaces/config"
	//	"database/sql"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	sqlConn *sqlx.DB
	queries *PreparedQueries
}

var Conn *DB

func init() {
	var err error
	Conn = &DB{}

	Conn.sqlConn, err = sqlx.Open("mysql", config.DbStr)
	if err != nil {
		log.Fatal("Cannot init mysql err=", err)
	}
	Conn.initQueries()

}
