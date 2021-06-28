package database

import (
	"log"

	"github.com/gobuffalo/pop/v5"
)

func PopDb() *pop.Connection {
	pop.Debug = true

	db, err := pop.Connect("production")

	if err != nil {
		log.Panic(err)
	}

	return db
}
