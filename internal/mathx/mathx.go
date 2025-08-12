package mathx

import (
	"errors"
	"math"
)

// IsPositiveAndEven demonstrates partial line coverage with chained conditions.
// This line is ideal for partial coverage: only one branch will execute in tests.
func IsPositiveAndEven(n int) bool {
	return n > 0 && isEven(n) && expensiveValidation(n) // Partial coverage target: chained && conditions
}

// ProcessNumbers shows partial coverage with multiple function calls on same line.
// Tests will only cover some execution paths, making this line partially covered.
func ProcessNumbers(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("negative numbers not allowed")
	}
	// Partial coverage target: only one of these functions will be called in some tests
	return complexCalculation(a, b) + simpleCalculation(a, b), validateResult(a, b) // Mixed execution paths
}

// SmartDivide demonstrates partial coverage with inline conditional logic.
// The ternary-like pattern using if creates opportunities for partial line coverage.
func SmartDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	// Partial coverage target: conditional execution on same line
	result := func() float64 { if math.Abs(b) < 1e-10 { return math.Inf(1) } else { return a / b } }() // Inline conditional
	return result, nil
}

// Helper functions to create realistic partial coverage scenarios

func isEven(n int) bool {
	return n%2 == 0
}

func expensiveValidation(n int) bool {
	// Simulate expensive operation that might be skipped in some test paths
	return n < 1000000
}

func complexCalculation(a, b int) int {
	return a*b + a + b
}

func simpleCalculation(a, b int) int {
	return a + b
}

func validateResult(a, b int) error {
	if a+b > 10000 {
		return errors.New("result too large")
	}
	return nil
}