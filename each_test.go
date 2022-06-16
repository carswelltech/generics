package generics

import (
	"strings"
	"testing"
)

func testEach[K comparable](t *testing.T, args ...K) {
	i := 0
	for a := range Each(args...) {
		if a != args[i] {
			t.Errorf("Each(%v) error, expected %v, got %v", args, args[i], a)
		}
		i++
	}
}
func TestEach(t *testing.T) {
	testEach[int](t, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	testEach[string](t, strings.Split("For now is the time for all good men to come to the aid of their country", " ")...)
	testEach[float32](t, 1.1, 1.2, 2.1, 2.2, 3.14159, 2.71, 3.0)
	testEach[rune](t, []rune("For now is the time for all good men to come to the aid of their country")...)
}
