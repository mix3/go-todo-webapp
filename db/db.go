package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mix3/go-todo-webapp/options"
)

type DB struct {
	dbmap *gorp.DbMap
}

func New(opts options.Options) (*DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		opts.DBUser,
		opts.DBPass,
		opts.DBHost,
		opts.DBPort,
		opts.DBName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	todoMap := dbmap.AddTableWithName(Todo{}, "todo").SetKeys(true, "Id")
	todoMap.ColMap("Text").SetNotNull(true)
	todoMap.ColMap("Done").SetNotNull(true)

	if err = dbmap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}

	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "todo:", log.Lmicroseconds))

	return &DB{dbmap: dbmap}, nil
}

func (db *DB) Close() {
	db.dbmap.Db.Close()
}
