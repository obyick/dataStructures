package sortingTypes

// MergeSort takes an array and sorts it using the MergeSort
func MergeSort(v []int) {
	if len(v) <= 1 { // if the array has length 1 or less, it is already sorted
		return
	}

	// Divide the array into two halves using two pointers instead of copying the array
	mid := len(v) / 2
	MergeSort(v[:mid])
	MergeSort(v[mid:])

	// Check if the last element of the first array is smaller than or equal to the first element of the second array
	if v[mid-1] <= v[mid] {
		return
	}

	// Merge the sorted halves into a single sorted array
	copy(v, Merge(v[:mid], v[mid:]))
}

// Takes two sorted arrays and merges them into a single sorted array
func Merge(v1 []int, v2 []int) []int {
	merged := make([]int, 0, len(v1)+len(v2))

	// Check if the last element of the first array is smaller or equal to the first element of the second array
	if v1[len(v1)-1] <= v2[0] {
		return append(merged, v1...) // Concatenate the two arrays without merging
	}

	// Create three pointers i, j, and k initialized to 0
	i, j, k := 0, 0, 0

	// Loop while both arrays have elements remaining
	for j < len(v1) && k < len(v2) {
		// If the current element of the first array is smaller than the current element of the second array, add it to the merged array and increment j
		if v1[j] < v2[k] {
			merged = append(merged, v1[j])
			j++
		} else { // Otherwise, add the current element of the second array to the merged array and increment k
			merged = append(merged, v2[k])
			k++
		}
		// Increment i
		i++
	}

	// Add the remaining elements of v1 to the merged array
	merged = append(merged, v1[j:]...)

	// Add the remaining elements of v2 to the merged array
	merged = append(merged, v2[k:]...)

	// Return the merged sorted array
	return merged
}
