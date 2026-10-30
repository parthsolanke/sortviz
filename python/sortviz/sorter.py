from __future__ import annotations

from collections.abc import Callable

from .helpers import SortStep, append_step, heapify, merge_sort_helper, quick_sort_helper


def selection_sort(values: list[int]) -> list[SortStep]:
    arr = list(values)
    steps: list[SortStep] = []
    n = len(arr)

    for i in range(n - 1):
        min_idx = i
        for j in range(i + 1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]

        append_step(steps, arr)

    return steps


def bubble_sort(values: list[int]) -> list[SortStep]:
    arr = list(values)
    steps: list[SortStep] = []
    n = len(arr)

    for i in range(n - 1):
        swapped = False
        for j in range(n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                swapped = True

        append_step(steps, arr)

        if not swapped:
            break

    return steps


def insertion_sort(values: list[int]) -> list[SortStep]:
    arr = list(values)
    steps: list[SortStep] = []
    n = len(arr)

    for i in range(1, n):
        key = arr[i]
        j = i - 1

        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key

        append_step(steps, arr)

    return steps


def merge_sort(values: list[int]) -> list[SortStep]:
    steps: list[SortStep] = []
    merge_sort_helper(list(values), steps)
    return steps


def quick_sort(values: list[int]) -> list[SortStep]:
    arr = list(values)
    steps: list[SortStep] = []
    quick_sort_helper(arr, 0, len(arr) - 1, steps)
    return steps


def shell_sort(values: list[int]) -> list[SortStep]:
    arr = list(values)
    steps: list[SortStep] = []
    n = len(arr)

    gap = n // 2
    while gap > 0:
        for i in range(gap, n):
            temp = arr[i]
            j = i
            while j >= gap and arr[j - gap] > temp:
                arr[j] = arr[j - gap]
                j -= gap
            arr[j] = temp

        append_step(steps, arr)
        gap //= 2

    return steps


def cocktail_sort(values: list[int]) -> list[SortStep]:
    arr = list(values)
    steps: list[SortStep] = []
    start = 0
    end = len(arr) - 1

    while start < end:
        swapped = False

        for i in range(start, end):
            if arr[i] > arr[i + 1]:
                arr[i], arr[i + 1] = arr[i + 1], arr[i]
                swapped = True
        append_step(steps, arr)

        end -= 1

        for i in range(end, start, -1):
            if arr[i - 1] > arr[i]:
                arr[i - 1], arr[i] = arr[i], arr[i - 1]
                swapped = True
        append_step(steps, arr)

        start += 1

        if not swapped:
            break

    return steps


def heap_sort(values: list[int]) -> list[SortStep]:
    arr = list(values)
    steps: list[SortStep] = []
    n = len(arr)

    for i in range(n // 2 - 1, -1, -1):
        heapify(arr, n, i, steps)

    for i in range(n - 1, 0, -1):
        arr[0], arr[i] = arr[i], arr[0]
        append_step(steps, arr)
        heapify(arr, i, 0, steps)

    return steps


ALGORITHMS: dict[str, Callable[[list[int]], list[SortStep]]] = {
    "bubble": bubble_sort,
    "insertion": insertion_sort,
    "selection": selection_sort,
    "merge": merge_sort,
    "quick": quick_sort,
    "heap": heap_sort,
    "shell": shell_sort,
    "cocktail": cocktail_sort,
}


def sort(values: list[int], algorithm: str) -> list[SortStep]:
    try:
        sorter = ALGORITHMS[algorithm]
    except KeyError as exc:
        raise ValueError(f"unknown sorting algorithm: {algorithm}") from exc

    return sorter(values)
