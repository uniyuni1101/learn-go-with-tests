package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/uniyuni1101/learn-go-with-tests/chapter22/websocket"
)

func TestRecodeingWinsAndRetrievingThem(t *testing.T) {
	db, cleandb := createTempFile(t, "[]")
	defer cleandb()

	store, err := poker.NewFileSystemStore(db)
	assertNoError(t, err)

	server, _ := poker.NewPlayerServer(store, dummyGame)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}

		assertLeague(t, got, want)
	})
}
