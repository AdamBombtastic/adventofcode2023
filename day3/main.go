package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func loadText(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	out, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(out)
}

type RawPartNumber struct {
	row     int
	column  int
	content string
}

func (r *RawPartNumber) Number() int {
	val, _ := strconv.Atoi(r.content)
	return val
}
func (r *RawPartNumber) IsAdjacent(row int, column int) bool {
	bound_left := r.column - 1
	bound_right := r.column + len(r.content)
	bound_top := r.row - 1
	bound_bottom := r.row + 1
	// fmt.Printf("Checking %d,%d against %d,%d,%d,%d\n", row, column, bound_top, bound_bottom, bound_left, bound_right)
	return row >= bound_top && row <= bound_bottom && column <= bound_right && column >= bound_left
}

type SymbolReference struct {
	row    int
	column int
	symbol string
}

func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

// Returns true if r is the following symbols: =,%,@,/,+,-,*,#,&,$
func IsSymbol(r rune) bool {
	return r == '=' || r == '%' || r == '@' || r == '/' || r == '+' || r == '-' || r == '*' || r == '#' || r == '&' || r == '$'
}

// To simplify for part 1 I'm just going to grab all the numbers outright.

func toRawPartNumbers(lines []string) ([]*RawPartNumber, []*SymbolReference) {
	parts := make([]*RawPartNumber, 0)
	symbols := make([]*SymbolReference, 0)
	for rowIndex, line := range lines {
		buffer := ""
		for colIndex, current := range line {
			if IsNumber(current) {
				buffer += string(current)
				continue
			}
			if IsSymbol(current) {
				symbols = append(symbols, &SymbolReference{row: rowIndex, column: colIndex, symbol: string(current)})
			}
			if len(buffer) > 0 {
				parts = append(parts, &RawPartNumber{content: buffer, row: rowIndex, column: colIndex - len(buffer)})
				buffer = ""
			}
		}
		if len(buffer) > 0 {
			parts = append(parts, &RawPartNumber{content: buffer, row: rowIndex, column: len(line) - len(buffer)})
			buffer = ""
		}
	}
	return parts, symbols
}

// isPartAdjacent returns true if the part is adjacent to the symbol.
func isPartAdjacent(part *RawPartNumber, symbol *SymbolReference) bool {
	return part.IsAdjacent(symbol.row, symbol.column)
}

func SolvePart2(rawNumbers []*RawPartNumber, rawSymbols []*SymbolReference) int {
	sum := 0
	for _, rawSymbol := range rawSymbols {
		if rawSymbol.symbol != "*" {
			continue
		}
		adjacentNumbers := make([]*RawPartNumber, 0)
		for _, rawNumber := range rawNumbers {
			if isPartAdjacent(rawNumber, rawSymbol) {
				adjacentNumbers = append(adjacentNumbers, rawNumber)
			}
		}
		if len(adjacentNumbers) == 2 {
			sum += (adjacentNumbers[0].Number() * adjacentNumbers[1].Number())
		}
	}
	return sum
}

func SolvePart1(rawNumbers []*RawPartNumber, rawSymbols []*SymbolReference) int {
	sum := 0
	symbolMap := make(map[string]bool)
	for _, rawNumber := range rawNumbers {
		for _, rawSymbol := range rawSymbols {
			symbolMap[rawSymbol.symbol] = true
			if isPartAdjacent(rawNumber, rawSymbol) {
				// fmt.Printf("Found %d adjacent to %s\n", rawNumber.Number(), rawSymbol.symbol)
				sum += rawNumber.Number()
				break
			}
		}
	}
	// fmt.Printf("len rawNumbers: %d\n", len(rawNumbers))
	// fmt.Printf("len rawSymbols: %d\n", len(rawSymbols))
	// for symbol := range symbolMap {
	// 	fmt.Printf("Symbol: |%s|\n", symbol)
	// }
	return sum
}

func main() {
	input := loadText(os.Args[1])
	sum := 0
	rawNumbers, rawSymbols := toRawPartNumbers(strings.Split(input, "\n"))
	sum = SolvePart1(rawNumbers, rawSymbols)
	println("Part 1 ans: ", sum)

	sum = SolvePart2(rawNumbers, rawSymbols)
	println("Part 2 ans: ", sum)
}
