package main

import (
	"log"
	"my_lib_log/internal/db"
	"my_lib_log/internal/pkg/app"
)

func main() {

	db, err := db.InitPostgresDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	a, err := app.New(db)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}

}
