package main

import (
	"log"
	"net/http"

	poker "github.com/uniyuni1101/learn-go-with-tests/chapter20/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
