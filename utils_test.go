package utils

import (
	"reflect"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("do a thing", func(t *testing.T) {
		elements := []string{"a", "b", "a", "c", "d", "b", "a"}
		got := Counter(elements)
		want := map[string]int{
			"a": 3,
			"b": 2,
			"c": 1,
			"d": 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestSplitLinesToInts(t *testing.T) {
	t.Run("split to ints", func(t *testing.T) {
		line := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
		got := SplitLinesToInts(line)
		want := [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestAbs(t *testing.T) {
	t.Run("converts negative to positive", func(t *testing.T) {
		got := Abs(-10)
		want := 10

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("keeps positive the same", func(t *testing.T) {
		got := Abs(10)
		want := 10

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
