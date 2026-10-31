from __future__ import annotations

from dataclasses import dataclass, field


@dataclass
class SortStep:
    array: list[int] = field(default_factory=list)

    def to_dict(self) -> dict[str, list[int]]:
        return {"array": list(self.array)}


def append_step(steps: list[SortStep], arr: list[int]) -> list[SortStep]:
    if not steps or steps[-1].array != arr:
        steps.append(SortStep(array=list(arr)))
    return steps


def merge_sort_helper(arr: list[int], steps: list[SortStep]) -> list[int]:
    if len(arr) <= 1:
        return arr

    mid = len(arr) // 2
    left = merge_sort_helper(arr[:mid], steps)
    right = merge_sort_helper(arr[mid:], steps)

    merged = merge(left, right)

    append_step(steps, merged)
    return merged


def merge(left: list[int], right: list[int]) -> list[int]:
    result: list[int] = []
    i = 0
    j = 0

    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1

    result.extend(left[i:])
    result.extend(right[j:])

    return result


def partition(arr: list[int], low: int, high: int, steps: list[SortStep]) -> int:
    pivot = arr[high]
    i = low - 1

    for j in range(low, high):
        if arr[j] < pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]

    arr[i + 1], arr[high] = arr[high], arr[i + 1]

    append_step(steps, arr)
    return i + 1


def quick_sort_helper(arr: list[int], low: int, high: int, steps: list[SortStep]) -> None:
    if low < high:
        pi = partition(arr, low, high, steps)
        quick_sort_helper(arr, low, pi - 1, steps)
        quick_sort_helper(arr, pi + 1, high, steps)


def heapify(arr: list[int], n: int, i: int, steps: list[SortStep]) -> None:
    largest = i
    left = 2 * i + 1
    right = 2 * i + 2

    if left < n and arr[left] > arr[largest]:
        largest = left
    if right < n and arr[right] > arr[largest]:
        largest = right

    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]
        append_step(steps, arr)
        heapify(arr, n, largest, steps)
