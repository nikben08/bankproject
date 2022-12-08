package database

import (
	"bankproject/models"
	"bankproject/seeds"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	var err error // define error here to prevent overshadowing the global DB

	dbURL := "postgres://postgres:08112001@localhost:5432/bank"

	DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Migrator().DropTable(models.User{}, models.Bank{}, models.Interest{})

	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(models.User{}, models.Bank{}, models.Interest{})

	if err != nil {
		log.Fatal(err)
	}

	superUser := seeds.SuperUser()
	accessLevel, _ := strconv.Atoi(superUser["accessLevel"])
	var admin = &models.User{
		Username:    superUser["username"],
		Hash:        superUser["hash"],
		AccessLevel: accessLevel}

	if result := DB.Create(&admin); result.Error != nil {
		fmt.Println("Couldn't create super user")
	}

	return DB
}
