package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanishId = "Spanish"
const frenchId = "French"
const exclamation = "!"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + exclamation
}

func greetingPrefix(language string) (helloPrefix string) {
	switch language {
	case spanishId:
		helloPrefix = spanishHelloPrefix
	case frenchId:
		helloPrefix = frenchHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Zbyszek", ""))
}
