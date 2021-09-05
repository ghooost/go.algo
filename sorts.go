package main

import (
	"fmt"

	"example.com/bubble"
	"example.com/merge"
	"example.com/quick"
)

func main() {
	fmt.Println("Bubble sort")
	arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	fmt.Println(arr)
	bubble.Bubble(arr[:])
	fmt.Println(arr)

	fmt.Println("Merge sort")
	arr = []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	fmt.Println(arr)
	merge.Merge(arr[:])
	fmt.Println(arr)

	fmt.Println("Quick sort")
	arr = []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	fmt.Println(arr)
	quick.Quick(arr[:])
	fmt.Println(arr)

}
