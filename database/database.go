package database

import (
	"context"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type Db struct {
	Client driver.Database
}

var db *Db

func Get() (*Db, error) {
	if db == nil {
		connection, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{"http://localhost:8529"},
		})
		if err != nil {
			return nil, err
		}

		client, err := driver.NewClient(driver.ClientConfig{
			Connection: connection,
			Authentication: driver.BasicAuthentication("root", "rootpassword"),
		})
		if err != nil {
			return nil, err
		}

		database, err := client.Database(context.Background(), "databse name goes here")
		if err != nil {
			return nil, err
		}

		db = &Db{
			Client: database,
		}
	}

	return db, nil
}