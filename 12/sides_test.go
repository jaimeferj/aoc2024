package main

import (
	"reflect"
	"testing"
)

func TestCalculateRegionSides(t *testing.T) {
	bounds := [2]int{4, 4}
	tests := []struct {
		name     string
		grid     [][]bool
		expected int
	}{
		{
			name: "Single cell region",
			grid: [][]bool{
				{true, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			expected: 4, // Single cell region.
		},
		{
			name: "L-shaped region",
			grid: [][]bool{
				{true, false, false, false},
				{true, true, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			expected: 6,
		},
		{
			name: "Irregular region",
			grid: [][]bool{
				{true, false, false, false},
				{true, true, false, false},
				{false, true, true, false},
				{false, false, false, false},
			},
			expected: 10,
		},
		{
			name: "Central Space",
			grid: [][]bool{
				{true, true, true, false},
				{true, false, true, false},
				{true, true, true, false},
				{false, false, false, false},
			},
			expected: 8,
		},
		{
			name: "Spike",
			grid: [][]bool{
				{false, true, false, false},
				{true, true, true, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			expected: 8,
		},
		{
			name: "Spike2",
			grid: [][]bool{
				{true, true, true, false},
				{false, true, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			expected: 8,
		},
		{
			name: "TouchingBorders",
			grid: [][]bool{
				{true, true, true, true},
				{true, true, false, true},
				{true, false, true, true},
				{true, true, true, true},
			},
			expected: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			regionVertex := make([][]bool, bounds[0]+1)
			for i := range bounds[0] + 1 {
				regionVertex[i] = make([]bool, bounds[1]+1)
			}
			for i, row := range test.grid {
				for j, value := range row {
					if value {

						regionVertex[i][j] = true
						regionVertex[i][j+1] = true
						regionVertex[i+1][j] = true
						regionVertex[i+1][j+1] = true
					}
				}
			}
			sides := calcSides(regionVertex, test.grid, bounds)
			if !reflect.DeepEqual(sides, test.expected) {
				t.Errorf("Got %v, expected %v", sides, test.expected)
			}
		})
	}
}
