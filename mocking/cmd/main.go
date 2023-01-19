package main

import (
	"os"
	"time"

	"github.com/hereisjohnny2/mocking"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
