package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"seeder/seeds"
)

func main() {
	dbConn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	for _, seed := range seeds.All() {
		if err := seed.Run(dbConn); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
