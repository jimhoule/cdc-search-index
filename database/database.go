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
			Endpoints: []string{
				os.Getenv("DB_URL"), os.Getenv("DB_PORT"),
			},
		})
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		client, err := driver.NewClient(driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		})
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		database, err := client.Database(context.Background(), os.Getenv("DB_NAME"))
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
