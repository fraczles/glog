package db

import (
	"database/sql"
	"github.com/fraczles/glog/shared/models"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func InitDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	db, err := sql.Open("sqlite3", "./blog.db")
	checkErr(err, "Create tables failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(models.Post{}, "posts").SetKeys(true, "ID")

	// create table if doesn't exist yet
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}
