package main

import (
	"go_practice/gorm_src/database"
	"go_practice/gorm_src/models"
	"log"
)

func main() {
	dbUser, dbPassword, dbName := "root", "root", "gorm_crud_example"

	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.DB().Ping()

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Contact{})

	defer db.Close()

	route := SetupRoutes(db)

	route.Run(":8000")
}
