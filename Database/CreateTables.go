package Database

import (
	"log"
)

//unused code prepared (not finished) to check if table exist and then create one

func (db *Database) CreateTables() error {
	if err := db.createNewTable("clubssss"); err != nil {
		return err
	}

	return nil
}

func (db *Database) checkIfTableExist(tableName string) bool {
	var result int

	db.Database.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = ?;", tableName).Scan(&result)

	if result == 0 {
		return false
	}

	return true
}

func (db *Database) createNewTable(tableName string) error {
	if exist := db.checkIfTableExist(tableName); exist {
		log.Println("Table already exist")
		return nil
	}

	log.Println("Creating new table")
	return nil
}
