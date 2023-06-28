package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/uniyuni1101/learn-go-with-tests/chapter21/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type \"{Name} wins\" to record a win")

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)

	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPocker()
}
