package main

import (
	"task-manager/internal/db"
	"task-manager/internal/db/adapters"
)

func main() {

	// Conectando no banco postgresql
	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.ConnectToDatabase(postgresAdapter)

	if err != nil {
		panic(err)
	}
}
