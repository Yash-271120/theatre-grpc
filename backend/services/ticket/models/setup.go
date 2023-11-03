package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host, hostexists := os.LookupEnv("DBHOST")
	port, portexists := os.LookupEnv("DBPORT")
	user, userexists := os.LookupEnv("DBUSER")
	password, passwordexists := os.LookupEnv("DBPASS")
	dbname, dbnameexists := os.LookupEnv("DBNAME")

	if !hostexists || !portexists || !userexists || !passwordexists || !dbnameexists {
		log.Fatalf("Error loading environment variables for DB")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to DB : %v", err)
	}

	err = db.AutoMigrate(&Ticket{})
	if err != nil {
		log.Fatalf("Error migrating DB : %v", err)
	}

	log.Println("Connected to DB successfully")
	DB = db

}
