package storage

import (
	"context"
	"github.com/gaizkadc/wallet-master-6000/config"
	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
)

var DB *pg.DB

func SetupDB() {
	configuration := config.GetConfig()

	database := configuration.Database.Dbname
	host := configuration.Database.Host
	port := configuration.Database.Port
	username := configuration.Database.Username
	password := configuration.Database.Password

	DB = pg.Connect(&pg.Options{
		Addr:     host + ":" + port,
		User:     username,
		Password: password,
		Database: database,
	})

	// Check if database is connected
	ctx := context.Background()

	if err := DB.Ping(ctx); err != nil {
		panic(err)
	}
}
