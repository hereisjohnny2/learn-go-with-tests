package mocking

import (
	"fmt"
	"io"
	"time"
)

const final_word = "Go!"
const count = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{
		duration,
		sleep,
	}
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := count; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, final_word)

}
