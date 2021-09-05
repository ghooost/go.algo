package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bubble sort")
	arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	fmt.Println(arr)
	Bubble(arr[:])
	fmt.Println(arr)

	fmt.Println("Merge sort")
	arr = []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	fmt.Println(arr)
	Merge(arr[:])
	fmt.Println(arr)

	fmt.Println("Quick sort")
	arr = []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	fmt.Println(arr)
	Quick(arr[:])
	fmt.Println(arr)
}
