package sort

func bubble_sort(array []int) []int {
	var count int = len(array)
	for i := 0; i < count; i++ {
		for j := 0; j < count - 1; j++ {
			if array[j] > array [j + 1] {
				array[j], array[j + 1] = array[j + 1], array[j]
			}
		}
	}
	return array
}
