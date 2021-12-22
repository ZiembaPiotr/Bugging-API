package Database

import (
	"backend-BD/Entities"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Database struct {
	Database *gorm.DB
}

func SetUpDBConnection() *gorm.DB {
	var psqlInfo string
	var err error
	var db Database

	if psqlInfo, err = generatePsqlInfo(); err != nil {
		log.Fatal(err)
	}

	if db.Database, err = gorm.Open(postgres.Open(psqlInfo)); err != nil {
		log.Fatal(err)
	}

	if err := db.autoMigration(); err != nil {
		log.Fatal(err)
	}

	return db.Database
}

func generatePsqlInfo() (string, error) {

	var environmentVariables [6]string
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	environmentVariables[0] = os.Getenv("DIALECT")
	environmentVariables[1] = os.Getenv("HOST")
	environmentVariables[2] = os.Getenv("DBPORT")
	environmentVariables[3] = os.Getenv("DBNAME")
	environmentVariables[4] = os.Getenv("USER")
	environmentVariables[5] = os.Getenv("PASSWORD")

	if err = validEnvironmentVar(environmentVariables); err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		environmentVariables[1], environmentVariables[2], environmentVariables[4], environmentVariables[3], environmentVariables[5])

	return psqlInfo, nil
}

func validEnvironmentVar(environmentVariables [6]string) error {
	for i, envVar := range environmentVariables {
		if i != 5 && envVar == "" {
			return errors.New("unable to get environment variables")
		}
	}
	return nil
}

func (db *Database) autoMigration() error {
	if err := db.Database.AutoMigrate(&Entities.Guest{}); err != nil {
		return err
	}

	if err := db.Database.AutoMigrate(&Entities.HotelAdmin{}); err != nil {
		return err
	}

	if err := db.Database.AutoMigrate(&Entities.HotelDetails{}); err != nil {
		return err
	}

	if err := db.Database.AutoMigrate(&Entities.Reservation{}); err != nil {
		return err
	}

	if err := db.Database.AutoMigrate(&Entities.Rooms{}); err != nil {
		return err
	}
	return nil
}
