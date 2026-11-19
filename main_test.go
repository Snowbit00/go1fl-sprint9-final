package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "negative size", size: -1, want: 0},
		{name: "zero size", size: 0, want: 0},
		{name: "one element", size: 1, want: 1},
		{name: "multiple elements", size: 100, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.want {
				t.Fatalf("len(generateRandomElements(%d)) = %d, want %d", tt.size, len(data), tt.want)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive number", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{7}, want: 7},
		{name: "maximum first", data: []int{9, 1, 3}, want: 9},
		{name: "maximum last", data: []int{3, 1, 9}, want: 9},
		{name: "repeated maximum", data: []int{2, 9, 1, 9, 3}, want: 9},
		{name: "negative elements", data: []int{-8, -2, -5}, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{7}, want: 7},
		{name: "fewer elements than chunks", data: []int{-4, -2, -6}, want: -2},
		{name: "one element per chunk", data: []int{-9, -8, -7, -6, -5, -1, -3, -2}, want: -1},
		{name: "maximum in a middle chunk", data: []int{1, 2, 3, 4, 5, 6, 20, 8, 9, 10, 11, 12, 13, 14, 15, 16}, want: 20},
		{name: "maximum in remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 20}, want: 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
