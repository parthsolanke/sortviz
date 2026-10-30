import {
  appendStep,
  heapify,
  mergeSortHelper,
  quickSortHelper,
} from './helpers.js';

export function selectionSort(values) {
  const arr = [...values];
  const steps = [];
  const n = arr.length;

  for (let i = 0; i < n - 1; i += 1) {
    let minIdx = i;
    for (let j = i + 1; j < n; j += 1) {
      if (arr[j] < arr[minIdx]) {
        minIdx = j;
      }
    }
    [arr[i], arr[minIdx]] = [arr[minIdx], arr[i]];

    appendStep(steps, arr);
  }

  return steps;
}

export function bubbleSort(values) {
  const arr = [...values];
  const steps = [];
  const n = arr.length;

  for (let i = 0; i < n - 1; i += 1) {
    let swapped = false;
    for (let j = 0; j < n - i - 1; j += 1) {
      if (arr[j] > arr[j + 1]) {
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
        swapped = true;
      }
    }

    appendStep(steps, arr);

    if (!swapped) {
      break;
    }
  }

  return steps;
}

export function insertionSort(values) {
  const arr = [...values];
  const steps = [];
  const n = arr.length;

  for (let i = 1; i < n; i += 1) {
    const key = arr[i];
    let j = i - 1;

    while (j >= 0 && arr[j] > key) {
      arr[j + 1] = arr[j];
      j -= 1;
    }
    arr[j + 1] = key;

    appendStep(steps, arr);
  }

  return steps;
}

export function mergeSort(values) {
  const steps = [];
  mergeSortHelper([...values], steps);
  return steps;
}

export function quickSort(values) {
  const arr = [...values];
  const steps = [];
  quickSortHelper(arr, 0, arr.length - 1, steps);
  return steps;
}

export function shellSort(values) {
  const arr = [...values];
  const steps = [];
  const n = arr.length;

  for (let gap = Math.floor(n / 2); gap > 0; gap = Math.floor(gap / 2)) {
    for (let i = gap; i < n; i += 1) {
      const temp = arr[i];
      let j = i;
      while (j >= gap && arr[j - gap] > temp) {
        arr[j] = arr[j - gap];
        j -= gap;
      }
      arr[j] = temp;
    }

    appendStep(steps, arr);
  }

  return steps;
}

export function cocktailSort(values) {
  const arr = [...values];
  const steps = [];
  let start = 0;
  let end = arr.length - 1;

  while (start < end) {
    let swapped = false;

    for (let i = start; i < end; i += 1) {
      if (arr[i] > arr[i + 1]) {
        [arr[i], arr[i + 1]] = [arr[i + 1], arr[i]];
        swapped = true;
      }
    }
    appendStep(steps, arr);

    end -= 1;

    for (let i = end; i > start; i -= 1) {
      if (arr[i - 1] > arr[i]) {
        [arr[i - 1], arr[i]] = [arr[i], arr[i - 1]];
        swapped = true;
      }
    }
    appendStep(steps, arr);

    start += 1;

    if (!swapped) {
      break;
    }
  }

  return steps;
}

export function heapSort(values) {
  const arr = [...values];
  const steps = [];
  const n = arr.length;

  for (let i = Math.floor(n / 2) - 1; i >= 0; i -= 1) {
    heapify(arr, n, i, steps);
  }

  for (let i = n - 1; i > 0; i -= 1) {
    [arr[0], arr[i]] = [arr[i], arr[0]];
    appendStep(steps, arr);
    heapify(arr, i, 0, steps);
  }

  return steps;
}

export const ALGORITHMS = {
  bubble: bubbleSort,
  insertion: insertionSort,
  selection: selectionSort,
  merge: mergeSort,
  quick: quickSort,
  heap: heapSort,
  shell: shellSort,
  cocktail: cocktailSort,
};

export function sort(values, algorithm) {
  const sorter = ALGORITHMS[algorithm];
  if (!sorter) {
    throw new Error(`unknown sorting algorithm: ${algorithm}`);
  }

  return sorter(values);
}
