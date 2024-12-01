package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		part := strings.Split(line, "   ")

		left[i], _ = strconv.Atoi(part[0])
		right[i], _ = strconv.Atoi(part[1])
	}

	// Part 1
	//sort.Ints(left)
	//sort.Ints(right)
	//sum := 0
	//for i := 0; i < len(lines); i++ {
	//	diff := left[i] - right[i]
	//	if diff < 0 {
	//		diff *= -1
	//	}
	//	sum += diff
	//}
	//fmt.Println(sum)

	// Part 2
	simSum := 0
	for i := 0; i < len(left); i++ {
		count := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				count++
			}
		}

		similarity := left[i] * count
		simSum += similarity
	}
	fmt.Println(simSum)
}
