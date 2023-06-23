package main

import (
	"os"
	"time"

	"github.com/uniyuni1101/learn-go-with-tests/chapter16/clockface/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
