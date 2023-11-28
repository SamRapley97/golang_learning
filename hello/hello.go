package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const russianHelloPrefix = "Privet, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name


	
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix

	case "Spanish":
		prefix = spanishHelloPrefix

	case "Russian":
		prefix = russianHelloPrefix
	
	default:
		prefix = englishHelloPrefix
	}
	return


}

func main() {
	fmt.Println(Hello("Sam", ""))
}
