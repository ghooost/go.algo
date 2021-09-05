package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
	arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	Merge(arr[:])
	wanted := []int{0, 1, 1, 2, 3, 4, 5, 6, 12}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func TestMergeEmpty(t *testing.T) {
	arr := []int{}
	Merge(arr[:])
	wanted := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func TestMergeSingle(t *testing.T) {
	arr := []int{1}
	Merge(arr[:])
	wanted := []int{1}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf(`%v not equal %v`, arr, wanted)
		}
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
		Merge(arr[:])
	}
}
