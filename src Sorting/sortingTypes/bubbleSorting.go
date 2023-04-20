package sortingTypes

func BubbleSort(v []int) {
	for i := 0; i < len(v)-1; i++ { // Loop through the array
		for j := 0; j < len(v)-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j] // Swap the elements
			}
		}
	}
}
