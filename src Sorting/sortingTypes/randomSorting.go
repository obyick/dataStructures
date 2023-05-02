package sortingTypes

import "math/rand" // Import the math/rand package for generating random numbers

// Takes an integer size as an argument and returns an integer slice
func RandomSort(size int) []int {
	// Create a new integer slice of size
	v := make([]int, size)

	// Loop through the slice and set each value to a random integer between 0 and 9 (inclusive)
	for i := 0; i < size; i++ {
		v[i] = rand.Intn(99)
	}

	// Return the filled slice
	return v
}
