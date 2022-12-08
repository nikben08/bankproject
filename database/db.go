package database

import (
	"bankproject/models"
	"bankproject/seeds"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	var err error // define error here to prevent overshadowing the global DB

	connStr := "postgres://postgres:08112001@localhost:5432/battlegame"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("232233223")
	db.Exec("DROP SCHEMA public CASCADE")
	db.Close()

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
