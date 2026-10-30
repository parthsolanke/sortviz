package sorter

type SortStep struct {
	Array []int `json:"array"`
}

func SelectionSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]

		steps = appendStep(steps, arr)
	}

	return steps
}

func BubbleSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		steps = appendStep(steps, arr)

		if !swapped {
			break
		}
	}

	return steps
}

func InsertionSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key

		steps = appendStep(steps, arr)
	}

	return steps
}

func MergeSort(arr []int) []SortStep {
	var steps []SortStep
	sortedArr := mergeSortHelper(arr, &steps)
	_ = sortedArr
	return steps
}

func QuickSort(arr []int) []SortStep {
	var steps []SortStep
	quickSortHelper(arr, 0, len(arr)-1, &steps)
	return steps
}

func ShellSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)

	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}

		steps = appendStep(steps, arr)
	}

	return steps
}

func HeapSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, &steps)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		steps = appendStep(steps, arr)
		heapify(arr, i, 0, &steps)
	}

	return steps
}
