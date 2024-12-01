package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	re := regexp.MustCompile(`\s+`)

	lines := strings.Split(input, "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		part := re.Split(lines[i], 2)

		l, _ := strconv.Atoi(part[0])
		left[i] = l
		r, _ := strconv.Atoi(part[1])
		right[i] = r
	}

	// sort.Ints(left)
	// sort.Ints(right)

	// sum := 0
	// for i := 0; i < len(lines); i++ {
	// 	diff := left[i] - right[i]
	// 	if diff < 0 {
	// 		diff *= -1
	// 	}
	// 	sum += diff
	// }
	// fmt.Println(sum)

	// iterate through the left array
	simSum := 0
	for i := 0; i < len(left); i++ {
		// Count the number of elements of left[i] in the right array
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
