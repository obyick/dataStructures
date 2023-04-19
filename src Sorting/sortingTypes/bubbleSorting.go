package sortingTypes

func BubbleSort(v []int) {
	// Loop through the array
	for i := 0; i < len(v)-1; i++ {
		for j := 0; j < len(v)-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j] // Swap the elements at index j and j + 1
			}
		}
	}
}
