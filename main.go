package main

import (
	"fmt"
	"log"
	"main/database"
	"main/router"
	"main/search"
	"main/searchclient"
	"main/users"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Panic(err)
	}

	db := database.Get()
	mainRouter := router.Get()
	searchClient := searchclient.Get()

	users.Init(mainRouter, db)
	search.Init(mainRouter, searchClient)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("HTTP_URL"), os.Getenv("HTTP_PORT")),
		Handler: mainRouter,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
