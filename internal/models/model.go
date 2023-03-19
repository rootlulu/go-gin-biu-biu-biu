package model

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/rootlulu/go-gin-biu-biu-biu/internal/config"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteIns struct {
	DataBaseSource string
}

var sqliteType = "sqlite3"

// Init ...
func Init() {
	initS := `
	CREATE TABLE IF NOT EXISTS lulu (id INTEGER PRIMARY KEY, name STRING UNIQUE, password STRING);
	INSERT INTO lulu(name, password) VALUES('ysm', 'Judi');
	INSERT INTO lulu(name, password) VALUES('lulu', 'Judi wife');
	`
	_, err := Exec(initS)
	if err != nil {
		logging.Fatal("Table create failed.")
	}
	logging.Info("Table lulu created!")
}

func SqliteContext(funcName string, arg string) []reflect.Value {
	db, err := sql.Open(config.DB.Type, config.DB.Path+config.DB.File)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()
	dbClass := reflect.ValueOf(db)
	method := dbClass.MethodByName(funcName)
	return method.Call([]reflect.Value{reflect.ValueOf(arg)})
}

func Query(query string) (*sql.Rows, error) {
	res := SqliteContext("Query", query)
	rowsV := res[0]
	errV := res[1]
	err, _ := errV.Interface().(error)
	return rowsV.Interface().(*sql.Rows), err
}

func QueryRow(query string) *sql.Row {
	row := SqliteContext("QueryRow", query)
	return row[0].Interface().(*sql.Row)
}

func Exec(query string) (sql.Result, error) {
	res := SqliteContext("Exec", query)
	resV := res[0]
	errV := res[1]
	// TODO 1: why the type assertion worked?
	// TODO 2: why var i interface{}; i.(interface{}) will fail?
	result, _ := resV.Interface().(sql.Result)
	return result, errV.Interface().(error)
}

type User struct {
	Username string
	Password string
}
