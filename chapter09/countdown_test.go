package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &CountdownOperationSpy{}

		CountDown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spy := &CountdownOperationSpy{}

		CountDown(spy, spy)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spy.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, spy.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
