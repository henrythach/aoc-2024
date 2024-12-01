package utils

func Counter[T comparable](elements []T) map[T]int {
	results := make(map[T]int)
	for _, element := range elements {
		results[element]++
	}
	return results
}
