package mathx

import (
	"testing"
)

func TestIsPositiveAndEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"positive even", 4, true},
		{"positive odd", 5, false},        // This will cause partial coverage - && chain breaks early
		{"negative even", -2, false},      // This will cause partial coverage - first condition fails
		{"zero", 0, false},                // This will cause partial coverage - first condition fails
		{"large positive even", 1000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPositiveAndEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsPositiveAndEven(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestProcessNumbers(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expectErr   bool
		expectedVal int
	}{
		{"positive small numbers", 2, 3, false, 16}, // 2*3+2+3 + 2+3 = 16
		{"negative a", -1, 5, true, 0},              // Error case - partial coverage
		{"negative b", 3, -2, true, 0},              // Error case - partial coverage
		{"zero values", 0, 0, false, 0},             // 0*0+0+0 + 0+0 = 0
		{"large numbers", 5000, 6000, true, 0},      // Will trigger validateResult error (5000+6000 > 10000)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ProcessNumbers(tt.a, tt.b)
			
			if tt.expectErr {
				if err == nil {
					t.Errorf("ProcessNumbers(%d, %d) expected error but got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("ProcessNumbers(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expectedVal {
					t.Errorf("ProcessNumbers(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expectedVal)
				}
			}
		})
	}
}

func TestSmartDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expectErr bool
		checkInf  bool
	}{
		{"normal division", 10.0, 2.0, false, false},
		{"division by zero", 5.0, 0.0, true, false},     // Error path - partial coverage
		{"very small divisor", 1.0, 1e-15, false, true}, // Will trigger inline conditional for Inf
		{"negative division", -8.0, 4.0, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SmartDivide(tt.a, tt.b)
			
			if tt.expectErr {
				if err == nil {
					t.Errorf("SmartDivide(%f, %f) expected error but got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("SmartDivide(%f, %f) unexpected error: %v", tt.a, tt.b, err)
				}
				if tt.checkInf {
					if result != result || result == 0 { // Check for Inf or unusual values
						// This is expected for very small divisors
					}
				}
			}
		})
	}
}

// Additional edge case tests to ensure we don't achieve 100% coverage
func TestPartialCoverageScenarios(t *testing.T) {
	t.Run("expensive validation not reached", func(t *testing.T) {
		// This test specifically avoids triggering expensiveValidation
		// by using odd numbers, creating partial line coverage
		result := IsPositiveAndEven(7)
		if result != false {
			t.Errorf("Expected false for odd number")
		}
	})

	t.Run("complex calculation only", func(t *testing.T) {
		// This test hits the line but doesn't trigger all execution paths
		// The validateResult might not be called in the same way
		result, _ := ProcessNumbers(1, 1)
		if result != 5 { // 1*1+1+1 + 1+1 = 5
			t.Errorf("Expected 5, got %d", result)
		}
	})
}

// Benchmark to show the functions work (optional, but good practice)
func BenchmarkIsPositiveAndEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPositiveAndEven(42)
	}
}

func BenchmarkProcessNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessNumbers(10, 20)
	}
}