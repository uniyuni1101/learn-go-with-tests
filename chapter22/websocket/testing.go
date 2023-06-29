package poker

import "testing"

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	v := s.Scores[player]
	return v
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := store.WinCalls[0]

	if got != winner {
		t.Errorf("didn't record correct winnner, got %q, want %q", got, winner)
	}
}
