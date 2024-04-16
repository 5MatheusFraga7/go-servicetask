package db

import "fmt"

type Database interface {
	Connect() error
	Close() error
}

func ConnectToDatabase(database Database) error {
	err := database.Connect()
	if err != nil {
		return err
	}

	fmt.Println("Conectamos!!!!!: %v", database)

	err = database.Close()
	if err != nil {
		return err
	}

	return nil
}
