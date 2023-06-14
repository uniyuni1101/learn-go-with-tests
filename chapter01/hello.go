package main

import "fmt"

const englishHelloPrefix = "Hello, "
const english = "English"

const japaneseHelloPrefix = "こんにちは、"
const japanese = "Japanese"

const frenchHelloPrefix = "Bonjour, "
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	var prefix string
	switch language {
	case english:
		prefix = englishHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
