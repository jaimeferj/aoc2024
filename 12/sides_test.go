package main

import (
	"reflect"
	"testing"
)

func TestCalculateRegionSides(t *testing.T) {
	bounds := [2]int{6, 6}
	tests := []struct {
		name     string
		grid     [][]bool
		expected int
	}{
		{
			name: "Single cell region",
			grid: [][]bool{
				{true, true, false, false, false, false},
				{true, true, false, false, false, false},
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
			},
			expected: 4, // Single cell region.
		},
		{
			name: "L-shaped region",
			grid: [][]bool{
				{true, true, false, false, false, false},
				{true, true, false, false, false, false},
				{true, true, true, true, false, false},
				{true, true, true, true, false, false},
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
			},
			expected: 6,
		},
		{
			name: "Irregular region",
			grid: [][]bool{
				{true, true, false, false, false, false},
				{true, true, false, false, false, false},
				{true, true, true, true, false, false},
				{true, true, true, true, false, false},
				{false, false, true, true, false, false},
				{false, false, true, true, false, false},
			},
			expected: 8,
		},
		{
			name: "Central Space",
			grid: [][]bool{
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
				{true, true, false, false, true, true},
				{true, true, false, false, true, true},
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
			},
			expected: 8,
		},
		{
			name: "Spike",
			grid: [][]bool{
				{false, false, true, true, false, false},
				{false, false, true, true, false, false},
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
			},
			expected: 8,
		},
		{
			name: "Spike2",
			grid: [][]bool{
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
				{false, false, true, true, false, false},
				{false, false, true, true, false, false},
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
