package sort

func quick_sort(array []int, start int, end int) {
	if start < end {
		i := start
		j := end
		base := array[i]

		for i < j {
			for i < j && array[j] >= base {
				j--
			}

			if i < j {
				array[i] = array[j]
				i++
			}

			for i < j && array[i] < base {
				i++
			}

			if i < j {
				array[j] = array[i]
				j--
			}
		}

		array[i] = base
		quick_sort(array, start, i - 1)
		quick_sort(array, i + 1, end)
	}
}
