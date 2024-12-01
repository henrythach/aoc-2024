package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`\s+`)

	file, err := os.ReadFile("cmd/day1/input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	input := string(file)

	lines := strings.Split(input, "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		part := re.Split(lines[i], 2)
		fmt.Println(part)

		l, _ := strconv.Atoi(part[0])
		left[i] = l
		r, _ := strconv.Atoi(part[1])
		right[i] = r
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(lines); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}
	fmt.Println(sum)
}
