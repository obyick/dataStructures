package sortingTypes

import "math/rand" // Import the math/rand package for generating random numbers

func RandomSort(size int) []int {
	v := make([]int, size) // Create a new integer slice of size
	for i := 0; i < size-1; i++ {
		v[i] = rand.Intn(10) // Set the value at index to a random integer
	}
	return v // Return the filled slice
}
