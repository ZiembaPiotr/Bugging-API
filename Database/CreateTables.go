package Database

import (
	"fmt"
	"gorm.io/gorm"
)

func CreateTables(db *gorm.DB) error {
	fmt.Println("Creating tables")

	return nil
}
