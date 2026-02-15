// package main declares an executable package.
package main

// fmt provides formatted I/O APIs such as Println.
import "fmt"

// Package-level constants used for language keys and greeting prefixes.
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a greeting string assembled from a prefix and name.
// Signature: two string parameters, one string return value.
func Hello(name string, language string) string {
	// Empty name fallback.
	if name == "" {
		name = "World"
	}

	// Function call expression: greetingPrefix(language) returns string.
	return greetingPrefix(language) + name
}

// greetingPrefix uses a named return value: `prefix string`.
// A bare `return` returns the current value of `prefix`.
func greetingPrefix(language string) (prefix string) {
	// switch selects one branch by comparing `language`.
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		// default runs when no case matches.
		prefix = englishHelloPrefix
	}
	return
}

// main is the program entry function.
func main() {
	fmt.Println(Hello("world", ""))
}
