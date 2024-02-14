package main

import "fmt"

const (
	french    = "French"
	spanish   = "Spanish"
	afrikaans = "Afrikaans"

	frenchHelloPrefix    = "Bonjour, "
	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	afrikaansHelloPrefix = "Goeiedag, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case afrikaans:
		prefix = afrikaansHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
