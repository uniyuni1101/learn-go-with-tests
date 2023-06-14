package iteration

import "testing"

func TestRepeart(t *testing.T) {
	repearted := Repeart("a")
	expected := "aaaaa"

	if repearted != expected {
		t.Errorf("expected %q but got %q", expected, repearted)
	}
}
