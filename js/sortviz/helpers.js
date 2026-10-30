export class SortStep {
  constructor(array = []) {
    this.array = [...array];
  }

  toDict() {
    return { array: [...this.array] };
  }

  toJSON() {
    return { array: [...this.array] };
  }
}

export function appendStep(steps, arr) {
  if (steps.length === 0 || !arraysEqual(steps[steps.length - 1].array, arr)) {
    steps.push(new SortStep(arr));
  }
  return steps;
}

export function arraysEqual(a, b) {
  return a.length === b.length && a.every((value, index) => value === b[index]);
}

export function mergeSortHelper(arr, steps) {
  if (arr.length <= 1) {
    return arr;
  }

  const mid = Math.floor(arr.length / 2);
  const left = mergeSortHelper(arr.slice(0, mid), steps);
  const right = mergeSortHelper(arr.slice(mid), steps);

  const merged = merge(left, right);

  appendStep(steps, merged);
  return merged;
}

export function merge(left, right) {
  const result = [];
  let i = 0;
  let j = 0;

  while (i < left.length && j < right.length) {
    if (left[i] < right[j]) {
      result.push(left[i]);
      i += 1;
    } else {
      result.push(right[j]);
      j += 1;
    }
  }

  result.push(...left.slice(i));
  result.push(...right.slice(j));

  return result;
}

export function partition(arr, low, high, steps) {
  const pivot = arr[high];
  let i = low - 1;

  for (let j = low; j < high; j += 1) {
    if (arr[j] < pivot) {
      i += 1;
      [arr[i], arr[j]] = [arr[j], arr[i]];
    }
  }

  [arr[i + 1], arr[high]] = [arr[high], arr[i + 1]];

  appendStep(steps, arr);
  return i + 1;
}

export function quickSortHelper(arr, low, high, steps) {
  if (low < high) {
    const pi = partition(arr, low, high, steps);
    quickSortHelper(arr, low, pi - 1, steps);
    quickSortHelper(arr, pi + 1, high, steps);
  }
}

export function heapify(arr, n, i, steps) {
  let largest = i;
  const left = 2 * i + 1;
  const right = 2 * i + 2;

  if (left < n && arr[left] > arr[largest]) {
    largest = left;
  }
  if (right < n && arr[right] > arr[largest]) {
    largest = right;
  }

  if (largest !== i) {
    [arr[i], arr[largest]] = [arr[largest], arr[i]];
    appendStep(steps, arr);
    heapify(arr, n, largest, steps);
  }
}
