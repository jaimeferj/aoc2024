package main

import "testing"
import "math"

func concatNumbers(a int, b int) int {
	valueDigits := int(math.Ceil(math.Log10(float64(b) + 0.1)))
	decimalPlaces := int(math.Pow10(valueDigits))
	return a*decimalPlaces + b
}

func concatNumbers2(a int, b int) int {
	concatProduct := 10
	for {
		if concatProduct <= b {
			concatProduct *= 10
		} else {
			break
		}
	}
	return a*concatProduct + b
}

func TestConcatNumber(t *testing.T) {
	// Define test cases
	tests := []struct {
		acc      int
		value    int
		expected int
	}{
		// Single-digit numbers
		{1, 1, 11},
		{2, 3, 23},
		{5, 7, 57},
		{9, 4, 94},

		// Multi-digit numbers
		{12, 34, 1234},
		{56, 78, 5678},
		{99, 99, 9999},

		// Large numbers
		{123, 456, 123456},
		{987, 654, 987654},

		// Combinations of single and multi-digit numbers
		{1, 10, 110},
		{10, 1, 101},
		{10, 10, 1010},
		{100, 1, 1001},
		{1, 100, 1100},

		// Edge cases
		{1, 999999999, 1999999999},
		{999999999, 1, 9999999991},
		{12345, 67890, 1234567890},
		{999, 1000, 9991000},

		// Extreme edge cases
		{999999, 999999, 999999999999},
		{1, 999999, 1999999},
		{999999, 1, 9999991},

		// Numbers with leading zero effect
		{1000, 1, 10001},
		{1, 1000, 11000},

		// Power of tens
		{10, 100, 10100},
		{100, 10, 10010},
		{1000, 100, 1000100},

		// Random pairs
		{123, 456, 123456},
		{789, 123, 789123},
		{1, 23456, 123456},
		{12345, 678, 12345678},

		// Very small and very large combinations
		{1, 1000000, 11000000},
		{1000000, 1, 10000001},
		{123456789, 987654321, 123456789987654321},
		{987654321, 123456789, 987654321123456789},

		// Random large numbers
		{111111, 222222, 111111222222},
		{999999, 123456, 999999123456},
		{123456, 999999, 123456999999},

		// Stress test for edge values near int32 limits
		{214748364, 7, 2147483647},
		{21474836, 47, 2147483647},
		{2, 147483647, 2147483647},

		// Random combinations with varying digit counts
		{1, 2, 12},
		{12, 3, 123},
		{123, 4, 1234},
		{1234, 5, 12345},
		{12345, 6, 123456},
		{123456, 7, 1234567},
		{1234567, 8, 12345678},
		{12345678, 9, 123456789},

		// Mixed combinations
		{111, 222, 111222},
		{222, 111, 222111},
		{123, 321, 123321},
		{321, 123, 321123},
		{100, 1000, 1001000},

		// Edge cases at boundaries of common ranges
		{99, 100, 99100},
		{100, 99, 10099},
		{999, 1000, 9991000},
		{1000, 999, 1000999},

		// Very large numbers near 1e9
		{999999999, 999999999, 999999999999999999},
		{123456789, 987654321, 123456789987654321},
		{987654321, 123456789, 987654321123456789},

		// Combinations with powers of 10
		{10, 100, 10100},
		{100, 10, 10010},
		{1000, 1, 10001},

		// Close to overflow (int64 testing)
		{922337203, 685477580, 922337203685477580},
		{1, 922337203, 1922337203},
		{922337203, 1, 9223372031},

		// Small edge cases
		{1, 1, 11},
		{2, 2, 22},
		{3, 3, 33},
		{4, 4, 44},
		{5, 5, 55},

		// More random numbers
		{123, 456, 123456},
		{789, 123, 789123},
		{1, 23456, 123456},
		{12345, 678, 12345678},
		{1, 23456789, 123456789},
		{23456789, 1, 234567891},

		{123456789, 987654321, 123456789987654321}, // Expected result, test against potential precision errors
		{1, 10, 110},     // Single digit and a power of 10
		{10, 1, 101},     // Power of 10 followed by a single digit
		{100, 10, 10010}, // Powers of 10 in both positions
		{1, 1, 11},       // Single digit concatenation
		{1000, 1, 10001}, // Multiple digits followed by 1:w
		{922337203, 685477580, 922337203685477580}, // Test within int64 range
		{922337203, 685477581, 922337203685477581}, // Expected to fail (overflow)
		{999999, 1, 9999991},                       // Large number followed by single digit
		{1, 999999, 1999999},                       // Single digit followed by large number
		{12345, 678, 12345678},                     // Medium numbers concatenation

	}
	// Run tests
	for _, test := range tests {
		result := concatNumbers(test.acc, test.value)
		if result != test.expected {
			t.Errorf("ConcatNumbers(%d, %d) = %d; want %d", test.acc, test.value, result, test.expected)
		}
	}
}
func TestConcatNumber2(t *testing.T) {
	// Define test cases
	tests := []struct {
		acc      int
		value    int
		expected int
	}{
		// Single-digit numbers
		{1, 1, 11},
		{2, 3, 23},
		{5, 7, 57},
		{9, 4, 94},

		// Multi-digit numbers
		{12, 34, 1234},
		{56, 78, 5678},
		{99, 99, 9999},

		// Large numbers
		{123, 456, 123456},
		{987, 654, 987654},

		// Combinations of single and multi-digit numbers
		{1, 10, 110},
		{10, 1, 101},
		{10, 10, 1010},
		{100, 1, 1001},
		{1, 100, 1100},

		// Edge cases
		{1, 999999999, 1999999999},
		{999999999, 1, 9999999991},
		{12345, 67890, 1234567890},
		{999, 1000, 9991000},

		// Extreme edge cases
		{999999, 999999, 999999999999},
		{1, 999999, 1999999},
		{999999, 1, 9999991},

		// Numbers with leading zero effect
		{1000, 1, 10001},
		{1, 1000, 11000},

		// Power of tens
		{10, 100, 10100},
		{100, 10, 10010},
		{1000, 100, 1000100},

		// Random pairs
		{123, 456, 123456},
		{789, 123, 789123},
		{1, 23456, 123456},
		{12345, 678, 12345678},

		// Very small and very large combinations
		{1, 1000000, 11000000},
		{1000000, 1, 10000001},
		{123456789, 987654321, 123456789987654321},
		{987654321, 123456789, 987654321123456789},

		// Random large numbers
		{111111, 222222, 111111222222},
		{999999, 123456, 999999123456},
		{123456, 999999, 123456999999},

		// Stress test for edge values near int32 limits
		{214748364, 7, 2147483647},
		{21474836, 47, 2147483647},
		{2, 147483647, 2147483647},

		// Random combinations with varying digit counts
		{1, 2, 12},
		{12, 3, 123},
		{123, 4, 1234},
		{1234, 5, 12345},
		{12345, 6, 123456},
		{123456, 7, 1234567},
		{1234567, 8, 12345678},
		{12345678, 9, 123456789},

		// Mixed combinations
		{111, 222, 111222},
		{222, 111, 222111},
		{123, 321, 123321},
		{321, 123, 321123},
		{100, 1000, 1001000},

		// Edge cases at boundaries of common ranges
		{99, 100, 99100},
		{100, 99, 10099},
		{999, 1000, 9991000},
		{1000, 999, 1000999},

		// Very large numbers near 1e9
		{999999999, 999999999, 999999999999999999},
		{123456789, 987654321, 123456789987654321},
		{987654321, 123456789, 987654321123456789},

		// Combinations with powers of 10
		{10, 100, 10100},
		{100, 10, 10010},
		{1000, 1, 10001},

		// Close to overflow (int64 testing)
		{922337203, 685477580, 922337203685477580},
		{1, 922337203, 1922337203},
		{922337203, 1, 9223372031},

		// Small edge cases
		{1, 1, 11},
		{2, 2, 22},
		{3, 3, 33},
		{4, 4, 44},
		{5, 5, 55},

		// More random numbers
		{123, 456, 123456},
		{789, 123, 789123},
		{1, 23456, 123456},
		{12345, 678, 12345678},
		{1, 23456789, 123456789},
		{23456789, 1, 234567891},

		{123456789, 987654321, 123456789987654321}, // Expected result, test against potential precision errors
		{1, 10, 110},     // Single digit and a power of 10
		{10, 1, 101},     // Power of 10 followed by a single digit
		{100, 10, 10010}, // Powers of 10 in both positions
		{1, 1, 11},       // Single digit concatenation
		{1000, 1, 10001}, // Multiple digits followed by 1:w
		{922337203, 685477580, 922337203685477580}, // Test within int64 range
		{922337203, 685477581, 922337203685477581}, // Expected to fail (overflow)
		{999999, 1, 9999991},                       // Large number followed by single digit
		{1, 999999, 1999999},                       // Single digit followed by large number
		{12345, 678, 12345678},                     // Medium numbers concatenation

	}
	// Run tests
	for _, test := range tests {
		result := concatNumbers2(test.acc, test.value)
		if result != test.expected {
			t.Errorf("ConcatNumbers(%d, %d) = %d; want %d", test.acc, test.value, result, test.expected)
		}
	}
}
