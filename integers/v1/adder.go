// package integers defines a reusable library package (not executable main).
package integers

// Add takes two int parameters and returns one int result.
// `x, y int` is Go's grouped parameter type syntax.
func Add(x, y int) int {
	// Arithmetic `+` for ints.
	return x + y
}
