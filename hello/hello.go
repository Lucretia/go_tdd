package main

import "fmt"

const englishHellPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHellPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
