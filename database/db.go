package database

import (
	"bankproject/config"
	"bankproject/models"
	"bankproject/seeds"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Config("DBHost"), config.Config("DBUsername"), config.Config("DBUserPassword"), config.Config("DBPort"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Exec("DROP DATABASE IF EXISTS bank;").Error; err != nil {
		panic(err)
	}

	if err := DB.Exec("CREATE DATABASE bank").Error; err != nil {
		panic(err)
	}

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Config("DBHost"), config.Config("DBUsername"), config.Config("DBUserPassword"), config.Config("DBName"), config.Config("DBPort"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	user_migration, err := ioutil.ReadFile("/bankproject/database/sql/user_migration.sql")

	if err != nil {
		log.Fatal(err)
	}

	bank_migration, err := ioutil.ReadFile("/bankproject/database/sql/bank_migration.sql")

	if err != nil {
		log.Fatal(err)
	}

	interest_migration, err := ioutil.ReadFile("/bankproject/database/sql/interest_migration.sql")

	if err != nil {
		log.Fatal(err)
	}

	sql1 := string(user_migration)
	sql2 := string(bank_migration)
	sql3 := string(interest_migration)

	if err := DB.Exec(sql1).Error; err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Exec(sql2).Error; err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Exec(sql3).Error; err != nil {
		panic(err)
	}

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
