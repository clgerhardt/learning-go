package main

import (
	"fmt"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return gettingLanguagePrefix(language) + name
}

func gettingLanguagePrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}
