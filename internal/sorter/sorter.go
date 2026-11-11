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
			if j != i {
				arr[j] = temp
				steps = appendStep(steps, arr)
			}
		}
	}

	return steps
}

func CocktailSort(arr []int) []SortStep {
	var steps []SortStep
	start := 0
	end := len(arr) - 1
	swapped := true

	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
		steps = appendStep(steps, arr)

		swapped = false
		end--
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		steps = appendStep(steps, arr)

		start++
	}

	return steps
}

func GnomeSort(arr []int) []SortStep {
	var steps []SortStep
	i := 0

	for i < len(arr) {
		if i == 0 || arr[i] >= arr[i-1] {
			i++
			continue
		}

		arr[i], arr[i-1] = arr[i-1], arr[i]
		steps = appendStep(steps, arr)
		i--
	}

	return steps
}

func CombSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)
	gap := n
	swapped := true

	for gap > 1 || swapped {
		gap = gap * 10 / 13
		if gap < 1 {
			gap = 1
		}

		swapped = false
		for i := 0; i+gap < n; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
		}

		if swapped {
			steps = appendStep(steps, arr)
		}
	}

	return steps
}

func OddEvenSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)
	sorted := false

	for !sorted {
		sorted = true

		for i := 1; i < n-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}

		for i := 0; i < n-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}

		steps = appendStep(steps, arr)
	}

	return steps
}

func RadixSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)
	if n == 0 {
		return steps
	}

	minVal := arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
	}

	work := make([]int, n)
	for i, v := range arr {
		work[i] = v - minVal
	}

	maxVal := work[0]
	for _, v := range work {
		if v > maxVal {
			maxVal = v
		}
	}

	buffer := make([]int, n)
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		counts := make([]int, 10)
		for _, v := range work {
			counts[(v/exp)%10]++
		}
		for i := 1; i < 10; i++ {
			counts[i] += counts[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			digit := (work[i] / exp) % 10
			counts[digit]--
			buffer[counts[digit]] = work[i]
		}
		copy(work, buffer)

		for i := range arr {
			arr[i] = work[i] + minVal
		}
		steps = appendStep(steps, arr)
	}

	return steps
}

func BucketSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)
	if n == 0 {
		return steps
	}

	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	bucketCount := n
	if maxVal == minVal {
		bucketCount = 1
	}

	buckets := make([][]int, bucketCount)
	span := maxVal - minVal + 1
	for _, v := range arr {
		idx := (v - minVal) * bucketCount / span
		if idx >= bucketCount {
			idx = bucketCount - 1
		}
		buckets[idx] = append(buckets[idx], v)
	}

	for _, bucket := range buckets {
		for i := 1; i < len(bucket); i++ {
			key := bucket[i]
			j := i - 1
			for j >= 0 && bucket[j] > key {
				bucket[j+1] = bucket[j]
				j--
			}
			bucket[j+1] = key
		}
	}

	idx := 0
	for _, bucket := range buckets {
		for _, v := range bucket {
			arr[idx] = v
			idx++
			steps = appendStep(steps, arr)
		}
	}

	return steps
}

func PancakeSort(arr []int) []SortStep {
	var steps []SortStep
	n := len(arr)

	for size := n; size > 1; size-- {
		maxIdx := 0
		for i := 1; i < size; i++ {
			if arr[i] > arr[maxIdx] {
				maxIdx = i
			}
		}

		if maxIdx == size-1 {
			continue
		}

		if maxIdx != 0 {
			flipPrefix(arr, maxIdx)
			steps = appendStep(steps, arr)
		}

		flipPrefix(arr, size-1)
		steps = appendStep(steps, arr)
	}

	return steps
}

func flipPrefix(arr []int, end int) {
	for i, j := 0, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func StoogeSort(arr []int) []SortStep {
	var steps []SortStep
	stoogeSortHelper(arr, 0, len(arr)-1, &steps)
	return steps
}

func CountingSort(arr []int) []SortStep {
	var steps []SortStep
	if len(arr) == 0 {
		return steps
	}

	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	counts := make([]int, maxVal-minVal+1)
	for _, v := range arr {
		counts[v-minVal]++
	}

	idx := 0
	for v := minVal; v <= maxVal; v++ {
		for counts[v-minVal] > 0 {
			arr[idx] = v
			idx++
			counts[v-minVal]--
			steps = appendStep(steps, arr)
		}
	}

	return steps
}
