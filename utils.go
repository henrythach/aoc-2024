package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Counter[T comparable](elements []T) map[T]int {
	results := make(map[T]int)
	for _, element := range elements {
		results[element]++
	}
	return results
}

func SplitLinesToInts(str string) [][]int {
	var results [][]int

	for _, line := range strings.Split(str, "\n") {
		var row []int
		for _, part := range strings.Fields(line) {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error converting '%s' to int: %v\n", part, err)
				continue
			}

			row = append(row, num)
		}

		results = append(results, row)
	}

	return results
}
