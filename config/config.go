package config

import (
	"casatemperatura/constants"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDBSqlx() *sqlx.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", constants.DBHOST, constants.DBPORT, constants.DBUSERNAME, constants.DBPASSWORD, constants.DBNAME)

	// open database
	db, err := sqlx.Open(constants.DBTYPE, psqlconn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
