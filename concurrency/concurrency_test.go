package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "fake.url.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://github.com",
		"fake.url.com",
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	want := map[string]bool{
		"https://google.com": true,
		"https://github.com": true,
		"fake.url.com":       false,
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(1 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf("URL: %d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
