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
		{"positive odd", 5, false},        // This will cause partial coverage - && chain breaks early
		{"negative even", -2, false},      // This will cause partial coverage - first condition fails
		{"zero", 0, false},                // This will cause partial coverage - first condition fails
		// Removed test cases that would complete the coverage
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
		{"negative a", -1, 5, true, 0},              // Error case - partial coverage
		{"zero values", 0, 0, false, 0},             // 0*0+0+0 + 0+0 = 0
		// Removed test cases to create partial coverage
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
		{"division by zero", 5.0, 0.0, true, false},     // Error path - partial coverage
		{"normal division", 10.0, 2.0, false, false},
		// Removed test that would trigger the inline conditional
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

// Test new functions with minimal coverage
func TestMaxOfThree(t *testing.T) {
	// Only test one path, leaving others uncovered
	result := MaxOfThree(5, 3, 1)
	if result != 5 {
		t.Errorf("MaxOfThree(5, 3, 1) = %d, want 5", result)
	}
}

func TestCheckRange(t *testing.T) {
	// Only test two cases, leaving most branches uncovered
	tests := []struct {
		input    int
		expected string
	}{
		{-5, "negative"},
		{50, "medium"},
	}
	
	for _, tt := range tests {
		result := CheckRange(tt.input)
		if result != tt.expected {
			t.Errorf("CheckRange(%d) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}