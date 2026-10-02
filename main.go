// Command orb-fleet-federation serves the HTTP API described in README.md.
package main

import (
	"log"
	"os"

	"github.com/Bob-xisuke/orb-fleet-federation/internal/api"
	"github.com/Bob-xisuke/orb-fleet-federation/internal/store"
)

func main() {
	address := os.Getenv("ADDR")
	if address == "" {
		address = "127.0.0.1:8080"
	}
	databasePath := os.Getenv("DB_PATH")
	if databasePath == "" {
		databasePath = "orb-fleet-federation.db"
	}

	st, err := store.Open(databasePath)
	if err != nil {
		log.Fatalf("open store: %v", err)
	}
	defer st.Close()

	if err := api.NewRouter(st).Run(address); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
