package main

import (
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	Bubble(arr[:])
	wanted := []int{0, 1, 1, 2, 3, 4, 5, 6, 12}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func TestBubbleEmpty(t *testing.T) {
	arr := []int{}
	Bubble(arr[:])
	wanted := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func TestBubbleSingle(t *testing.T) {
	arr := []int{1}
	Bubble(arr[:])
	wanted := []int{1}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
		Bubble(arr[:])
	}
}
