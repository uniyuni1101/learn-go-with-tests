package poker

import (
	"io"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	f, clean := createTempFile(t, "123456")
	defer clean()

	tape := &tape{f}

	tape.Write([]byte("abc"))

	f.Seek(0, 0)
	newFileContents, _ := io.ReadAll(f)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
