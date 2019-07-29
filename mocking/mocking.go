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

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := coundownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}

	Countdown(os.Stdout, sleeper)
}
