package sortingTypes

// Declare a function BubbleSort that takes an integer slice v as input and sorts it in non-decreasing order using the bubble sort algorithm
func BubbleSort(v []int) {
	// Loop through the array from the first to the second to last element
	for i := 0; i < len(v)-1; i++ {
		for j := 0; j < len(v)-1; j++ {
			// If the current element is greater than the next element, swap them
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j] // Swap the elements
			}
		}
	}
}
