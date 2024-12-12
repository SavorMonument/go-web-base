package main

import (
	"log"

	"github.com/alexedwards/scs/gormstore"
	"github.com/alexedwards/scs/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"generic.com/internal/app"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./db_data/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sessionManager := scs.New()
	if sessionManager.Store, err = gormstore.New(db); err != nil {
		log.Fatal(err)
	}

	a := app.NewApplication(sessionManager, db)
	a.Start()

}
