package main

import (
	"os"
)

type tape struct {
	f *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.f.Truncate(0)
	t.f.Seek(0, 0)
	return t.f.Write(p)
}
