// package main means this file builds as an executable command.
package main

// fmt is imported to use Println in main().
import "fmt"

// const declarations define immutable package-level values.
// In Go, untyped string constants are inferred from string literals.
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello has two parameters (name, language) and one return value (string).
// `func Hello(name string, language string) string` is equivalent to
// `func Hello(name, language string) string` in Go syntax.
func Hello(name string, language string) string {
	// `if` condition must be boolean; parentheses are not required in Go.
	if name == "" {
		// Reassign parameter variable when an empty string is passed.
		name = "World"
	}

	// `:=` declares and initializes a local variable.
	prefix := englishHelloPrefix

	// switch compares `language` against each case value.
	// If no case matches, prefix stays as English.
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	// `+` concatenates strings.
	return prefix + name
}

// main is the entry point for `go run` / built binary execution.
func main() {
	fmt.Println(Hello("world", ""))
}
