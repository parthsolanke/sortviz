package sorter

func deepCopyArray(arr []int) []int {
	return append([]int(nil), arr...)
}

func arrayChanged(lastStep, currentStep []int) bool {
	for i := range lastStep {
		if lastStep[i] != currentStep[i] {
			return true
		}
	}
	return false
}

func appendStep(steps []SortStep, arr []int) []SortStep {
	if len(steps) == 0 || arrayChanged(steps[len(steps)-1].Array, arr) {
		steps = append(steps, SortStep{Array: deepCopyArray(arr)})
	}
	return steps
}

func mergeSortHelper(arr []int, steps *[]SortStep) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSortHelper(arr[:mid], steps)
	right := mergeSortHelper(arr[mid:], steps)

	merged := merge(left, right)

	*steps = appendStep(*steps, merged)
	return merged
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func partition(arr []int, low, high int, steps *[]SortStep) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	*steps = appendStep(*steps, arr)
	return i + 1
}

func quickSortHelper(arr []int, low, high int, steps *[]SortStep) {
	if low < high {
		pi := partition(arr, low, high, steps)
		quickSortHelper(arr, low, pi-1, steps)
		quickSortHelper(arr, pi+1, high, steps)
	}
}

func heapify(arr []int, n, i int, steps *[]SortStep) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		*steps = appendStep(*steps, arr)
		heapify(arr, n, largest, steps)
	}
}
