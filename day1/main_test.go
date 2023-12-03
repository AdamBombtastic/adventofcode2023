package main

import (
	"testing"
)

func Test_ToNumberList(t *testing.T) {
	type testcase struct {
		line    string
		numbers []int
	}
	tests := []testcase{
		{"one", []int{1}},
		{"one1two", []int{1, 1, 2}},
		{"one1two2three", []int{1, 1, 2, 2, 3}},
		{"foo2baz", []int{2}},
		{"eightwothree", []int{8, 2, 3}}, // The nasty case
	}
	for _, tc := range tests {
		gotNumbers := toNumberList(tc.line)
		if len(gotNumbers) != len(tc.numbers) {
			t.Errorf("toNumberList(%q) = %v, want %v", tc.line, gotNumbers, tc.numbers)
		}
		for i, n := range gotNumbers {
			if n != tc.numbers[i] {
				t.Errorf("toNumberList(%q) = %v, want %v", tc.line, gotNumbers, tc.numbers)
			}
		}
	}
}
