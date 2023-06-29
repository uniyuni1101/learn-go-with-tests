package main

import (
	"log"
	"net/http"

	poker "github.com/uniyuni1101/learn-go-with-tests/chapter22/websocket"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)

	if err != nil {
		log.Fatalf("cloud not create server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
