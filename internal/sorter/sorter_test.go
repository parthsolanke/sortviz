package sorter

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := BubbleSort(arr)
	result := steps[len(steps)-1].Array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := InsertionSort(arr)
	result := steps[len(steps)-1].Array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := SelectionSort(arr)
	result := steps[len(steps)-1].Array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := MergeSort(arr)
	lastStep := steps[len(steps)-1].Array

	if !reflect.DeepEqual(lastStep, expected) {
		t.Errorf("Expected %v, got %v", expected, lastStep)
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := HeapSort(arr)
	result := steps[len(steps)-1].Array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestShellSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := ShellSort(arr)
	result := steps[len(steps)-1].Array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCocktailSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := CocktailSort(arr)
	result := steps[len(steps)-1].Array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	steps := QuickSort(arr)
	result := steps[len(steps)-1].Array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
