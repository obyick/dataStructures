package sortingTypes

// Define a function that performs an insertion sort on an input slice of integers
func InsertionSort(v []int) {
	// Loop through the array, starting at the second element
	for i := 1; i < len(v); i++ {
		// Starting at the current element, loop backwards through the array as long as the current element is smaller than the element to its left
		for j := i; j > 0 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j] // Swap the current element with the element to its left
		}
	}
}
