package database

import (
	"bankproject/models"
	"bankproject/seeds"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:08112001@localhost:5432/bank"
	DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Exec("DROP SCHEMA IF EXISTS public CASCADE;").Error; err != nil {
		panic(err)
	}

	if err := DB.Exec("Create SCHEMA public;").Error; err != nil {
		panic(err)
	}

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
