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
)

func main() {
	db := database.Get()
	mainRouter := router.Get()
	searchClient := searchclient.Get()

	users.Init(mainRouter, db)
	search.Init(mainRouter, searchClient)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", "5000"),
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
