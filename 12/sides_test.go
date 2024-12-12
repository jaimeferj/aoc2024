package main

import (
	"reflect"
	"testing"
)

func TestCalculateRegionSides(t *testing.T) {
	bounds := [2]int{3, 3}
	tests := []struct {
		name     string
		grid     [][]bool
		expected int
	}{
		{
			name: "Single cell region",
			grid: [][]bool{
				{true, false, false},
				{false, false, false},
				{false, false, false},
			},
			expected: 4, // Single cell region.
		},
		{
			name: "Rectangle region",
			grid: [][]bool{
				{true, true, false},
				{true, true, false},
				{false, false, false},
			},
			expected: 4,
		},
		{
			name: "L-shaped region",
			grid: [][]bool{
				{true, true, false},
				{true, false, false},
				{true, false, false},
			},
			expected: 6,
		},
		{
			name: "Irregular region",
			grid: [][]bool{
				{true, true, false},
				{true, true, true},
				{false, true, true},
			},
			expected: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sides := calcSides(test.grid, bounds)
			if !reflect.DeepEqual(sides, test.expected) {
				t.Errorf("Got %v, expected %v", sides, test.expected)
			}
		})
	}
}
