package main

import (
	"fmt"

	"go-coverage-test/internal/mathx"
)

func main() {
	fmt.Println("Go Coverage Test Demo")
	
	// Test IsPositiveAndEven
	fmt.Println("\n--- Testing IsPositiveAndEven ---")
	testNumbers := []int{4, 7, -2, 0, 42}
	for _, num := range testNumbers {
		result := mathx.IsPositiveAndEven(num)
		fmt.Printf("IsPositiveAndEven(%d) = %t\n", num, result)
	}

	// Test ProcessNumbers
	fmt.Println("\n--- Testing ProcessNumbers ---")
	testPairs := []struct{ a, b int }{
		{2, 3},
		{-1, 5},
		{10, 20},
		{0, 1},
	}
	for _, pair := range testPairs {
		result, err := mathx.ProcessNumbers(pair.a, pair.b)
		if err != nil {
			fmt.Printf("ProcessNumbers(%d, %d) error: %v\n", pair.a, pair.b, err)
		} else {
			fmt.Printf("ProcessNumbers(%d, %d) = %d\n", pair.a, pair.b, result)
		}
	}

	// Test SmartDivide
	fmt.Println("\n--- Testing SmartDivide ---")
	testDivisions := []struct{ a, b float64 }{
		{10.0, 2.0},
		{5.0, 0.0},
		{1.0, 1e-15},
		{-8.0, 4.0},
	}
	for _, div := range testDivisions {
		result, err := mathx.SmartDivide(div.a, div.b)
		if err != nil {
			fmt.Printf("SmartDivide(%.1f, %g) error: %v\n", div.a, div.b, err)
		} else {
			fmt.Printf("SmartDivide(%.1f, %g) = %g\n", div.a, div.b, result)
		}
	}

	fmt.Println("\nDemo complete! Run 'make coverage' to see coverage report.")
}