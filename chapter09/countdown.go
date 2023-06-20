package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func CountDown(w io.Writer, sleeper Sleeper) {
	for count := 3; count > 0; count-- {
		fmt.Fprintln(w, count)
		sleeper.Sleep()
	}
	fmt.Fprintln(w, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
