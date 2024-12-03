package main

import (
	utils "aoc-2024"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	var levels = utils.SplitLinesToInts(input)
	safeCount := 0
	for _, level := range levels {
		isIncreasing := level[1] > level[0]
		isSafe := true
		for i := 1; i < len(level); i++ {
			if level[i] == level[i-1] {
				isSafe = false
				break
			}

			// All increasing or all decreasing
			if isIncreasing != (level[i] > level[i-1]) {
				//fmt.Printf("Level %d: A, %s\n", row, level)
				isSafe = false
				break
			}

			// Differ at least one and at most three
			diff := utils.Abs(level[i] - level[i-1])
			if diff > 3 {
				//fmt.Printf("Level %d: Unsafe because the diff was %d, %s\n", row, diff, level)
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount++
		}
	}
	fmt.Println("Safe count:", safeCount)
}
