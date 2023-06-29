package poker_test

import (
	"fmt"
	"testing"
	"time"

	poker "github.com/uniyuni1101/learn-go-with-tests/chapter22/websocket"
)

func TestGameStart(t *testing.T) {
	cases := []struct {
		players int
		alerts  []scheduledAlert
	}{
		{5,
			[]scheduledAlert{
				{at: 0 * time.Second, amount: 100},
				{at: 10 * time.Minute, amount: 200},
				{at: 20 * time.Minute, amount: 300},
				{at: 30 * time.Minute, amount: 400},
				{at: 40 * time.Minute, amount: 500},
				{at: 50 * time.Minute, amount: 600},
				{at: 60 * time.Minute, amount: 800},
				{at: 70 * time.Minute, amount: 1000},
				{at: 80 * time.Minute, amount: 2000},
				{at: 90 * time.Minute, amount: 4000},
				{at: 100 * time.Minute, amount: 8000},
			},
		},
		{7,
			[]scheduledAlert{
				{at: 0 * time.Second, amount: 100},
				{at: 12 * time.Minute, amount: 200},
				{at: 24 * time.Minute, amount: 300},
				{at: 36 * time.Minute, amount: 400},
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("schedules alerts on game start for %d players", c.players), func(t *testing.T) {
			alerter := &SpyBlindAlerter{}
			game := poker.NewTexasHoldem(alerter, dummyPlayerStore)

			game.Start(c.players, dummyStdOut)

			checkSchedulingCases(t, c.alerts, alerter)
		})
	}
}

func TestGameFinish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(t *testing.T, alerts []scheduledAlert, alerter *SpyBlindAlerter) {
	t.Helper()

	for i, want := range alerts {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}

			got := alerter.alerts[i]

			assertScheduledAlert(t, got, want)
		})
	}
}
