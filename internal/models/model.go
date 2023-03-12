package model

import (
	"database/sql"
	"log"

	"github.com/rootlulu/go-gin-biu-biu-biu/internal/config"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"

	_ "github.com/mattn/go-sqlite3"
)

// Init ...
func Init() {
	log.Println("Model initing...  log")
	logging.Info("Model initing... logging")
	db, err := sql.Open(config.DB.Type, config.DB.Path+config.DB.File)
	if err != nil {
		logging.Fatal(err)
	}
	defer db.Close()

	initS := `
	CREATE TABLE lulu(id INTEGER PRIMARY KEY, name TEXT, password STRING);
	INSERT INTO lulu(name, password) VALUES('ysm', 'Judi');
	INSERT INTO lulu(name, password) VALUES('lulu', 'Judi wife');
	`
	_, err = db.Exec(initS)
	if err != nil {
		logging.Fatal(err)
	}
	logging.Info("Table lulu created!")

}
