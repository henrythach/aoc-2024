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
