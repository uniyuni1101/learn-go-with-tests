package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"https://www.example.com/index.html",
		"https://www.google.com",
		"http://github.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"https://www.example.com/index.html": true,
		"https://www.google.com":             true,
		"http://github.com":                  true,
		"waat://furhurterwe.geds":            false,
	}

	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
