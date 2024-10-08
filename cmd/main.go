package main

import (
	"fmt"
	"log"

	"github.com/sandeshtamanq/jwt/cmd/api"
	"github.com/sandeshtamanq/jwt/database"
	"github.com/sandeshtamanq/jwt/entity"
)

func main() {

	dsn := "host=localhost user= password= dbname= port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	db, err := database.NewPostgresStorage(dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("database connected successfully")
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Task{})

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("server running...")
}
