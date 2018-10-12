package mysql

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbMap   map[string]*sql.DB
	dbMutex sync.Mutex
)

// ---------------------------------------------------------------------------------------------------------------------

func init() {
	dbMap = make(map[string]*sql.DB)
}

// Get tx conn
func GetTx(dataSource string) (*sql.Tx, error) {
	db, err := GetDB(dataSource)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Get db conn
func GetDB(dataSource string) (*sql.DB, error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if db, ok := dbMap[dataSource]; ok {
		return db, nil
	}

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(200)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	dbMap[dataSource] = db

	return db, nil
}

func FreeDB(dataSource string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if db, ok := dbMap[dataSource]; ok {
		db.Close()
		delete(dbMap, dataSource)
	}
}

func FreeAllDB() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	for dataSource, db := range dbMap {
		db.Close()
		delete(dbMap, dataSource)
	}
}
