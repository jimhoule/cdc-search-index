package database

import (
	"context"
	"fmt"
	"os"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type Db struct {
	Client driver.Database
}

var db *Db

func Get() *Db {
	if db == nil {
		connection, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{"http://localhost:8529"},
		})
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		client, err := driver.NewClient(driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication("root", "rootpassword"),
		})
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		database, err := client.Database(context.Background(), "cdc-search-index")
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		db = &Db{
			Client: database,
		}
	}

	return db
}
