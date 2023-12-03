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

type Game struct {
	index    int
	CubeSets []*CubeSet
}

func (g *Game) IsPossible(red, green, blue int) bool {
	for _, cubeSet := range g.CubeSets {
		// fmt.Printf("red: %d, blue: %d, green: %d\n", cubeSet.Red, cubeSet.Blue, cubeSet.Green)
		if !cubeSet.IsPossible(red, green, blue) {
			return false
		}
	}
	return true
}
func (g *Game) MinPossibleColors() *CubeSet {
	minPossibleSet := &CubeSet{
		Red:   0,
		Blue:  0,
		Green: 0,
	}
	for _, cubeSet := range g.CubeSets {
		if cubeSet.Red > minPossibleSet.Red {
			minPossibleSet.Red = cubeSet.Red
		}
		if cubeSet.Blue > minPossibleSet.Blue {
			minPossibleSet.Blue = cubeSet.Blue
		}
		if cubeSet.Green > minPossibleSet.Green {
			minPossibleSet.Green = cubeSet.Green
		}
	}
	return minPossibleSet
}

type CubeSet struct {
	Red   int
	Blue  int
	Green int
}

// for problem one
func (c CubeSet) IsPossible(red, green, blue int) bool {
	return c.Red <= red && c.Blue <= blue && c.Green <= green
}

// for problem two
func (c CubeSet) Power() int {
	return c.Red * c.Blue * c.Green
}

func toGame(line string) *Game {
	parts1 := strings.Split(line, ":")
	gameInfo := strings.Split(strings.TrimSpace(parts1[0]), " ")
	index, _ := strconv.ParseInt(gameInfo[1], 10, 64)
	game := &Game{
		index: int(index),
	}

	cubeSetParts := strings.Split(strings.TrimSpace(parts1[1]), ";")
	for _, cubeSetPart := range cubeSetParts {
		cubeSet := &CubeSet{}
		diceRecords := strings.Split(strings.TrimSpace(cubeSetPart), ",")
		for _, diceRecord := range diceRecords {
			colorInfo := strings.Split(strings.TrimSpace(diceRecord), " ")
			switch strings.TrimSpace(colorInfo[1]) {
			case "red":
				red, _ := strconv.ParseInt(colorInfo[0], 10, 64)
				cubeSet.Red = int(red)
			case "blue":
				blue, _ := strconv.ParseInt(colorInfo[0], 10, 64)
				cubeSet.Blue = int(blue)
			case "green":
				green, _ := strconv.ParseInt(colorInfo[0], 10, 64)
				cubeSet.Green = int(green)
			}
		}
		game.CubeSets = append(game.CubeSets, cubeSet)
	}
	return game
}

func solvePart1(games []*Game) int {
	sum := 0
	for _, game := range games {
		if game.IsPossible(12, 13, 14) {
			sum += game.index
		}
	}
	return sum
}

func solvePart2(games []*Game) int {
	sum := 0
	for _, game := range games {
		minPossibleSet := game.MinPossibleColors()
		sum += minPossibleSet.Power()
	}
	return sum
}

func main() {
	input := loadText(os.Args[1])
	sum := 0
	games := []*Game{}
	for _, line := range strings.Split(input, "\n") {
		games = append(games, toGame(line))
	}

	// part 1
	sum = solvePart1(games)
	println(sum)

	// part 2
	sum = solvePart2(games)
	println(sum)
}
