package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchhHelloPrefix = "Bonjour, "
const englishHellPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchhHelloPrefix + name
	}

	return englishHellPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
