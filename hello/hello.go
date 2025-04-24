package main

import (
	"fmt"
)

const (
	russian = "Russian"
	spanish = "Spanish"

	emptyName = "World"

	englishHelloPrefix = "Hello, "
	russianHelloPrefix = "Привет, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = emptyName
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case russian:
		prefix = russianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "English"))
}
