package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord     = "Go!"
	coundownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := coundownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	// for i := coundownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// }

	// for i := coundownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }

	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}

	Countdown(os.Stdout, sleeper)
}
