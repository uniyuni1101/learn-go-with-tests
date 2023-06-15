package iteration

import (
	"fmt"
	"testing"
)

func TestRepeart(t *testing.T) {
	repeartHelper := func(t testing.TB, char string, repeart int, expected string) {
		t.Helper()
		repearted := Repeart(char, repeart)
		if repearted != expected {
			t.Errorf("expected %q but got %q", expected, repearted)
		}
	}

	t.Run("5 repeart", func(t *testing.T) {
		char := "a"
		repeart := 5
		expected := "aaaaa"

		repeartHelper(t, char, repeart, expected)
	})

	t.Run("10 repeart", func(t *testing.T) {
		char := "a"
		repeart := 10
		expected := "aaaaaaaaaa"

		repeartHelper(t, char, repeart, expected)
	})
}

func BenchmarkRepeart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeart("a", 5)
	}
}

func ExampleRepeart() {
	repearted := Repeart("a", 5)
	fmt.Println(repearted)
	// Output: aaaaa
}
