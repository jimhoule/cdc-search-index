package main

import (
	"fmt"
	"log"
	"main/router"
	"main/search"
	"main/searchclient"
	"main/users"
	"net/http"
)

func main() {
	searchClient := searchclient.Get()
	mainRouter := router.Get()

	users.Init(mainRouter)
	search.Init(searchClient)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", "3000"),
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}