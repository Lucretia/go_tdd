package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord     = "Go!"
	coundownStart = 3
)

func Countdown(out io.Writer) {
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
