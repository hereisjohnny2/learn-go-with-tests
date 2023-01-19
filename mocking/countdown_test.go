package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (sco *SpyCountdownOperations) Sleep() {
	sco.Calls = append(sco.Calls, sleep)
}

func (sco *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	sco.Calls = append(sco.Calls, write)
	return
}

type SpyDuration struct {
	durationSlept time.Duration
}

func (sd *SpyDuration) Sleep(duration time.Duration) {
	sd.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("should countdown from 3 to go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SpyCountdownOperations{}

		Countdown(buffer, sleeperSpy)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should call 'sleep' after every count", func(t *testing.T) {
		sco := &SpyCountdownOperations{}

		Countdown(sco, sco)

		got := sco.Calls
		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got calls %v, wanted %v", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyDuration := &SpyDuration{}

	sleeper := ConfigurableSleeper{sleepTime, spyDuration.Sleep}
	sleeper.Sleep()

	if spyDuration.durationSlept != sleepTime {
		t.Errorf("should have slept for %v, but got %v", sleepTime, spyDuration.durationSlept)
	}
}
