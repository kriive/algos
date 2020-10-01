package sort

// Bubble is a function that implements Bubble sort
// https://en.wikipedia.org/wiki/Bubble_sort
// Bubble sort is an educational sorting algorithm
// because it does not perform really well in real life
func Bubble(slice []int) []int {
	lastElement := len(slice) - 1

	for {
		for i := 0; i < lastElement; i++ {
			if slice[i] > slice[i+1] {
				swap(&slice[i], &slice[i+1])
			}
		}

		// We limit cycles, because we're sure that
		// every time we have ordered at least the last n
		// elements, so there's no need to check those
		if lastElement += -1; lastElement == 0 {
			break
		}
	}
	return slice
}

func swap(first *int, second *int) {
	*first, *second = *second, *first
}
