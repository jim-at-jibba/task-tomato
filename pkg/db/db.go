package db

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

// var db *sql.DB

type taskDB struct {
	DB      *sql.DB
	dataDir string
}

// openDB opens a SQLite database and stores that database in our special spot.
func OpenDB(path string) (*taskDB, error) {
	var dbName string
	if viper.GetBool("debug") {
		dbName = "task-dev.db"
	} else {
		dbName = "task.db"
	}
	db, err := sql.Open("sqlite3", filepath.Join(path, dbName))
	if err != nil {
		return nil, err
	}
	t := taskDB{db, path}
	if !t.tableExists("tasks") {
		err := t.createTable()
		if err != nil {
			return nil, err
		}
	}

	return &t, nil
}

func (t *taskDB) tableExists(name string) bool {
	if _, err := t.DB.Query("SELECT * FROM tasks"); err == nil {
		return true
	}
	return false
}

func (t *taskDB) createTable() error {
	_, err := t.DB.Exec(`CREATE TABLE "tasks" ( "id" INTEGER, "name" TEXT NOT NULL, "project" TEXT, "status" TEXT, "created" DATETIME, PRIMARY KEY("id" AUTOINCREMENT))`)
	return err
}
