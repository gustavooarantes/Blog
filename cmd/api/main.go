package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
