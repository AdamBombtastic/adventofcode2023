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

type Card struct {
	winning []string
	losing  []string
	index   int

	// dirty cache
	parsedWinMap map[string]bool
}

func (c *Card) winningNumbers() []int {
	if c.parsedWinMap == nil {
		c.parsedWinMap = make(map[string]bool)
		for _, s := range c.winning {
			c.parsedWinMap[s] = true
		}
	}

	out := make([]int, 0)
	for _, s := range c.losing {
		if c.parsedWinMap[s] {
			out = append(out, mustInt(s))
		}
	}
	return out
}

func mustInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}
func filterBlanks(ss []string) []string {
	out := make([]string, 0, len(ss))
	for _, s := range ss {
		if strings.TrimSpace(s) != "" {
			out = append(out, s)
		}
	}
	return out
}
func NewFromLine(line string) *Card {
	lineParts := strings.Split(line, ":")

	cardIndex := mustInt(filterBlanks(strings.Split(lineParts[0], " "))[1])
	numberParts := strings.Split(strings.TrimSpace(lineParts[1]), "|")

	winningNumberStrings := filterBlanks(strings.Split(strings.TrimSpace(numberParts[0]), " "))
	losingNumberStrings := filterBlanks(strings.Split(strings.TrimSpace(numberParts[1]), " "))

	return &Card{
		winning: winningNumberStrings,
		losing:  losingNumberStrings,
		index:   cardIndex,
	}
}

func solve1(cards []*Card) int {
	sum := 0
	for _, card := range cards {
		winList := card.winningNumbers()
		// fmt.Printf("%d: %v\n", card.index, winList)
		if len(winList) == 0 {
			continue
		}
		currentVal := 0
		for i := 0; i < len(winList); i++ {
			if currentVal == 0 {
				currentVal = 1
			} else {
				currentVal *= 2
			}
		}
		sum += currentVal
	}
	return sum
}

// This is a very slow solution, but it works -- I could optimize, but I won't.
func solve2(cards []*Card) int {
	called := make(map[int]int)
	toCall := make(map[int]int)
	// init
	fmt.Println(len(cards))
	for _, card := range cards {
		toCall[card.index] = 1
		called[card.index] = 0
	}

	index := 0
	for {
		if index > len(cards)-1 {
			break
		}
		current := cards[index]
		nextAmount := len(current.winningNumbers())
		for i := index + 1; i < index+1+nextAmount; i++ {
			toCall[i] = toCall[i] + 1
		}
		toCall[index] = toCall[index] - 1
		called[index] = called[index] + 1
		if toCall[index] <= 0 {
			index += 1
		}
	}
	sum := 0
	for k, v := range called {
		fmt.Printf("%d: %d\n", k, v)
		sum += v
	}
	return sum
}

func main() {
	input := loadText(os.Args[1])
	cards := make([]*Card, 0)
	for _, line := range strings.Split(input, "\n") {
		cards = append(cards, NewFromLine(line))
	}

	sum := solve1(cards)
	println(sum)

	// This solution is very slow, but it works -- I could optimize, but I won't.
	sum = solve2(cards)
	println(sum)
}
