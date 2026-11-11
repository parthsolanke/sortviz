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

func stoogeSortHelper(arr []int, low, high int, steps *[]SortStep) {
	if low >= high {
		return
	}

	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
		*steps = appendStep(*steps, arr)
	}

	if high-low+1 > 2 {
		third := (high - low + 1) / 3
		stoogeSortHelper(arr, low, high-third, steps)
		stoogeSortHelper(arr, low+third, high, steps)
		stoogeSortHelper(arr, low, high-third, steps)
	}
}

func insertionSortRange(arr []int, start, end int, steps *[]SortStep) {
	for i := start + 1; i < end; i++ {
		key := arr[i]
		j := i - 1
		for j >= start && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		*steps = appendStep(*steps, arr)
	}
}

func timMerge(arr []int, left, mid, right int, steps *[]SortStep) {
	leftPart := append([]int(nil), arr[left:mid]...)
	rightPart := append([]int(nil), arr[mid:right]...)

	i, j, k := 0, 0, left
	for i < len(leftPart) && j < len(rightPart) {
		if leftPart[i] <= rightPart[j] {
			arr[k] = leftPart[i]
			i++
		} else {
			arr[k] = rightPart[j]
			j++
		}
		k++
	}
	for i < len(leftPart) {
		arr[k] = leftPart[i]
		i++
		k++
	}
	for j < len(rightPart) {
		arr[k] = rightPart[j]
		j++
		k++
	}

	*steps = appendStep(*steps, arr)
}

func bitonicSortHelper(arr []int, low, cnt int, dir bool, steps *[]SortStep) {
	if cnt <= 1 {
		return
	}

	k := cnt / 2
	bitonicSortHelper(arr, low, k, !dir, steps)
	bitonicSortHelper(arr, low+k, cnt-k, dir, steps)
	bitonicMerge(arr, low, cnt, dir, steps)
}

func bitonicMerge(arr []int, low, cnt int, dir bool, steps *[]SortStep) {
	if cnt <= 1 {
		return
	}

	k := greatestPowerOfTwoLessThan(cnt)
	for i := low; i < low+cnt-k; i++ {
		if (arr[i] > arr[i+k]) == dir {
			arr[i], arr[i+k] = arr[i+k], arr[i]
			*steps = appendStep(*steps, arr)
		}
	}

	bitonicMerge(arr, low, k, dir, steps)
	bitonicMerge(arr, low+k, cnt-k, dir, steps)
}

func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k < n {
		k <<= 1
	}
	return k >> 1
}

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func insertTreeNode(root *treeNode, value int) *treeNode {
	if root == nil {
		return &treeNode{value: value}
	}
	if value < root.value {
		root.left = insertTreeNode(root.left, value)
	} else {
		root.right = insertTreeNode(root.right, value)
	}
	return root
}

func traverseInOrder(node *treeNode, visit func(int)) {
	if node == nil {
		return
	}
	traverseInOrder(node.left, visit)
	visit(node.value)
	traverseInOrder(node.right, visit)
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
