package main

import (
	"fmt"

	"./sortingTypes"
)

func main() {
	// Random array
	v := sortingTypes.RandomSort(30)
	fmt.Printf("Random Array: \t%v\n", v)

	//Bubble sorting
	sortingTypes.BubbleSort(v)
	fmt.Printf("Bubble sort: \t%v\n", v)

	// Random array
	v = sortingTypes.RandomSort(10)
	fmt.Printf("Random Array: \t%v\n", v)

	//select sorting
	sortingTypes.SelectSort(v)
	fmt.Printf("Select sort: %v\n", v)
}
