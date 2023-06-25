package main

import (
	"log"
	"net/http"
	"sync"
)

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func NewInMemoeyPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return i.store[player]
}

func (i *InMemoryPlayerStore) RecordWin(player string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[player]++
}

func main() {
	store := NewInMemoeyPlayerStore()
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
