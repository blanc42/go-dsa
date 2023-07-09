package sort

func SelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		min_index := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}

		arr[i], arr[min_index] = arr[min_index], arr[i]
	}
	return arr
}
