package main

import (
	utils "aoc-2024"
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
	leftCount := utils.Counter(left)
	rightCount := utils.Counter(right)
	for k, v := range leftCount {
		simSum += k * v * rightCount[k]
	}
	fmt.Println(simSum)
}
