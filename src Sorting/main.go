package main

import (
	"fmt"

	"./sortingTypes"
)

func main() {
	// Array size
	size := 10

	// Random array
	v := sortingTypes.RandomSort(size)
	fmt.Printf("Random Array: \t%v\n", v)

	//Bubble sorting
	sortingTypes.BubbleSort(v)
	fmt.Printf("Bubble Sort: \t%v\n", v)

	// Random array
	v = sortingTypes.RandomSort(size)
	fmt.Printf("Random Array: \t%v\n", v)

	//select sorting
	sortingTypes.SelectionSort(v)
	fmt.Printf("Selection Sort: %v\n", v)

	// Random array
	v = sortingTypes.RandomSort(size)
	fmt.Printf("Random Array: \t%v\n", v)

	// Insertion sorting
	sortingTypes.InsertionSort(v)
	fmt.Printf("Insertion Sort: %v\n", v)
}
