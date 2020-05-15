package sort

// BubbleSort - sorts slice of int using bubble sort algorythm
func BubbleSort(elements []int) {
	sorting := true
	for sorting {
		sorting = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] < elements[i+1] {
				sorting = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}
