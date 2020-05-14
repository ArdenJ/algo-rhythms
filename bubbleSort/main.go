package main

import (
	"fmt"
)

var sortThis []int = []int{1, 2, 3, 5, 7, 4, 0, 9}

func bubbleSort(input []int) {

	swapped := true

	for swapped {
		swapped = false

		for i := 1; i < len(sortThis); i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	fmt.Println(input)
}

func main() {
	bubbleSort(sortThis)
}
