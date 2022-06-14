package generics

import (
	`testing`
)

func TestZero(t *testing.T) {
	var i int
	var s string
	var f float32

	if i != Zero[int]() {
		t.Errorf("Zero(): expected %v, got %v",i,Zero[int]())
	}
	if s != Zero[string]() {
		t.Errorf("Zero(): expected %v, got %v",s,Zero[string]())
	}
	if f != Zero[float32]() {
		t.Errorf("Zero(): expected %v, got %v",f,Zero[float32]())
	}
}
