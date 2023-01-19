package http_race

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	durationA := measureURLRequestDuration(urlA)
	durationB := measureURLRequestDuration(urlB)

	if durationA > durationB {
		return urlB
	}

	return urlA
}

func SyncRacer(urlA, urlB string) (string, error) {
	return ConfigurableRacer(urlA, urlB, 10*time.Second)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func measureURLRequestDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
