package main

import (
	"fmt"
	"log"
	"main/router"
	"main/users"
	"net/http"
)

func main() {
	mainRouter := router.Get()

	users.Init(mainRouter)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", "3000"),
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}