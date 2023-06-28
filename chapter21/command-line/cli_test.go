package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/uniyuni1101/learn-go-with-tests/chapter21/command-line"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

type GameSpy struct {
	StartedWith  int
	StartCalled  bool
	FinishedWith string
	FinishCalled bool
}

func (g *GameSpy) Start(players int) {
	g.StartCalled = true
	g.StartedWith = players
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishedWith = winner
}

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func userSends(inputs ...string) *bytes.Buffer {
	buf := &bytes.Buffer{}
	for _, input := range inputs {
		fmt.Fprintln(buf, input)
	}
	return buf
}

func assertGameStartedWith(t *testing.T, game *GameSpy, players int) {
	t.Helper()

	if game.StartedWith != players {
		t.Errorf("got %d, expected number of player is %d", game.StartedWith, players)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Error("called game start, also expected not call game start")
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	if !game.FinishCalled {
		t.Error("not called game finish")
	}

	if game.FinishedWith != winner {
		t.Errorf("got %q, want winner is %q", game.FinishedWith, winner)
	}
}

func assertMessagesSentToUser(t *testing.T, out *bytes.Buffer, wantMessages ...string) {
	t.Helper()

	want := strings.Join(wantMessages, "")
	got := out.String()

	if got != want {
		t.Errorf("got %q, sent to stdout buy expected %+v", got, wantMessages)
	}
}

func assertScheduledAlert(t *testing.T, got, want scheduledAlert) {
	t.Helper()

	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("get scheduled time of %v, want %v", got.at, want.at)
	}
}
