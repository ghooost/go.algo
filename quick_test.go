package main

import (
	"testing"
)

func TestQuick(t *testing.T) {
	arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	Quick(arr[:])
	wanted := []int{0, 1, 1, 2, 3, 4, 5, 6, 12}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func TestQuickEmpty(t *testing.T) {
	arr := []int{}
	Quick(arr[:])
	wanted := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func TestQuickSingle(t *testing.T) {
	arr := []int{1}
	Quick(arr[:])
	wanted := []int{1}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
		Quick(arr[:])
	}
}
