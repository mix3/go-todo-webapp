package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mix3/go-todo-webapp/apps"
	"github.com/mix3/go-todo-webapp/db"
	"github.com/mix3/go-todo-webapp/options"
)

func main() {
	opts, err := options.Get()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(opts)
	if err != nil {
		log.Fatal(err)
	}

	app := apps.New(opts, db)
	defer app.Close()

	addr := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	log.Println("[main] starting...")
	log.Println("[main] running on", addr, "...")

	log.Fatal(http.ListenAndServe(addr, app))
}
