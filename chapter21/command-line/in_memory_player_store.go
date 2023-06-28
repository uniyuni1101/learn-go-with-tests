package poker

import "sync"

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

func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
