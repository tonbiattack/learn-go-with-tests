// package integers keeps tests in the same package, so unexported identifiers
// would also be accessible from this file.
package integers

// testing package provides *testing.T and test runner integrations.
import "testing"

// Test function names must start with `Test` and accept `*testing.T`.
func TestAdder(t *testing.T) {
	// `:=` declares local variables with inferred types.
	sum := Add(2, 2)
	expected := 4

	// `!=` compares values; when true, the test is marked as failed by Errorf.
	if sum != expected {
		// `%d` formats integers in base-10.
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
