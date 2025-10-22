package main

import "log"

// Entry point of the app
func main() {
	cfg := config{
		addr: ":8080",
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
