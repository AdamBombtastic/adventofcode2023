package main

import (
	"fmt"
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
func isNumber(b rune) bool {
	return b >= '0' && b <= '9'
}

var numbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// the issue is that the string 1eightwofoothree5 should produce [1,8,2,3,5], not [8,3]
func toNumberList(line string) []int {
	out := []int{}
	for i := 0; i < len(line); i++ {
		current := line[i]
		if isNumber(rune(current)) {
			out = append(out, int(current-'0'))
			continue
		}

		// not a number, look ahead -- the key here is managing the overlaps
		val := lookAhead(line, i)
		if val != -1 {
			out = append(out, val)
		}

	}
	return out
}

// lookAhead returns the number if it finds one, or -1 if it doesn't
func lookAhead(line string, i int) int {
	buffer := ""
	for j := i; j < len(line); j++ {
		if isNumber(rune(line[j])) {
			return -1
		}
		buffer += string(line[j])
		if val, ok := numbers[buffer]; ok {
			return val
		}
	}
	return -1
}

func parseLine(line string) int {
	numbers := toNumberList(line)
	if len(numbers) == 0 {
		panic("no numbers found")
	}
	val, _ := strconv.ParseInt(fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1]), 10, 64)
	return int(val)
}

// solves part 2 -- part 1 is just a matter of changing the parseLine function.
func main() {
	input := loadText(os.Args[1])
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += parseLine(line)
	}
	println(sum)
}
