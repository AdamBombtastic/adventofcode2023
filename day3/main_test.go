package main

import (
	"strings"
	"testing"
)

func Test_ParseRawNumbers(t *testing.T) {
	type testcase struct {
		lines           []string
		expected        []*RawPartNumber
		expectedSymbols []*SymbolReference
	}

	cases := []testcase{
		{
			lines: []string{
				"123...114..12",
			},
			expected: []*RawPartNumber{
				{row: 0, column: 0, content: "123"},
				{row: 0, column: 6, content: "114"},
				{row: 0, column: 11, content: "12"},
			},
			expectedSymbols: []*SymbolReference{},
		},
		{
			lines: []string{
				"123...114..12",
				"*.....+....*.",
			},
			expected: []*RawPartNumber{
				{row: 0, column: 0, content: "123"},
				{row: 0, column: 6, content: "114"},
				{row: 0, column: 11, content: "12"},
			},
			expectedSymbols: []*SymbolReference{
				{row: 1, column: 0, symbol: "*"},
				{row: 1, column: 6, symbol: "+"},
				{row: 1, column: 11, symbol: "*"},
			},
		},
		{
			lines: []string{
				"857*653",
			},
			expected: []*RawPartNumber{
				{row: 0, column: 0, content: "857"},
				{row: 0, column: 4, content: "653"},
			},
			expectedSymbols: []*SymbolReference{
				{row: 0, column: 3, symbol: "*"},
			},
		},
	}
	for _, c := range cases {
		actual, actualSymbols := toRawPartNumbers(c.lines)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d raw part numbers, got %d", len(c.expected), len(actual))
		}
		for i, e := range c.expected {
			a := actual[i]
			if a.row != e.row {
				t.Errorf("Expected row %d, got %d", e.row, a.row)
			}
			if a.column != e.column {
				t.Errorf("Expected column %d, got %d", e.column, a.column)
			}
			if a.content != e.content {
				t.Errorf("Expected content %s, got %s", e.content, a.content)
			}
		}
		if len(actualSymbols) != len(c.expectedSymbols) {
			t.Errorf("Expected %d symbols, got %d", len(c.expectedSymbols), len(actualSymbols))
		}
		for i, e := range c.expectedSymbols {
			a := actualSymbols[i]
			if a.row != e.row {
				t.Errorf("Expected row %d, got %d", e.row, a.row)
			}
			if a.column != e.column {
				t.Errorf("Expected column %d, got %d", e.column, a.column)
			}
			if a.symbol != e.symbol {
				t.Errorf("Expected symbol %s, got %s", e.symbol, a.symbol)
			}
		}
	}
}

func Test_isPartAdjacent(t *testing.T) {
	type testcase struct {
		part     *RawPartNumber
		symbol   *SymbolReference
		expected bool
	}
	testcases := []testcase{
		{
			part:     &RawPartNumber{row: 0, column: 0, content: "123"},
			symbol:   &SymbolReference{row: 1, column: 0, symbol: "*"},
			expected: true,
		},
		{
			part:     &RawPartNumber{row: 0, column: 0, content: "123"},
			symbol:   &SymbolReference{row: 1, column: 1, symbol: "*"},
			expected: true,
		},
		{
			part:     &RawPartNumber{row: 0, column: 0, content: "123"},
			symbol:   &SymbolReference{row: 1, column: 2, symbol: "*"},
			expected: true,
		},
		{
			part:     &RawPartNumber{row: 0, column: 0, content: "123"}, //123
			symbol:   &SymbolReference{row: 1, column: 3, symbol: "*"},  //   * <--- Valid because it's adjacent to the part
			expected: true,
		},
		{
			part:     &RawPartNumber{row: 0, column: 0, content: "123"}, //123
			symbol:   &SymbolReference{row: 1, column: 4, symbol: "*"},  //    * <--- Invalid because it's not adjacent to the part
			expected: false,
		},
		{
			part:     &RawPartNumber{row: 1, column: 1, content: "123"},
			symbol:   &SymbolReference{row: 0, column: 0, symbol: "*"}, // * <--- Invalid because it's not adjacent to the part
			expected: true,
		},
		{
			part:     &RawPartNumber{row: 1, column: 1, content: "123"},
			symbol:   &SymbolReference{row: -1, column: 1, symbol: "*"}, //  * <--- Valid because it's adjacent to the part
			expected: false,
		},
	}
	for _, c := range testcases {
		actual := isPartAdjacent(c.part, c.symbol)
		if actual != c.expected {
			t.Errorf("Expected %t, got %t", c.expected, actual)
		}
	}
}

func Test_SolvePart1(t *testing.T) {
	type testcase struct {
		inputFile string
		expected  int
		lines     []string
	}
	cases := []testcase{
		{
			inputFile: "test.txt",
			expected:  4361,
			lines:     nil,
		},
		{
			inputFile: "web_test.txt",
			expected:  24,
			lines:     nil,
		},
		{
			lines: []string{
				"100...100..12",
				"*.....@....**",
				"............1",
			},
			expected: 213,
		},
	}
	for _, c := range cases {
		if c.lines == nil {
			input := loadText(c.inputFile)
			c.lines = strings.Split(input, "\n")
		}
		rawNum, rawSymb := toRawPartNumbers(c.lines)
		actual := SolvePart1(rawNum, rawSymb)
		if actual != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, actual)
		}
	}
}
