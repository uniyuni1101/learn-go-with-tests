package iteration

import "testing"

func TestRepeart(t *testing.T) {
	repearted := Repeart("a")
	expected := "aaaaa"

	if repearted != expected {
		t.Errorf("expected %q but got %q", expected, repearted)
	}
}


func BenchmarkRepeart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeart("a")
	}
}
