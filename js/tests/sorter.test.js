import assert from 'node:assert/strict';
import test from 'node:test';

import {
  ALGORITHMS,
  SortStep,
  bubbleSort,
  cocktailSort,
  heapSort,
  insertionSort,
  mergeSort,
  quickSort,
  selectionSort,
  shellSort,
  sort,
} from '../sortviz/index.js';

const SAMPLE = [5, 3, 8, 4, 2];
const EXPECTED = [2, 3, 4, 5, 8];

function assertSorts(sorter) {
  const steps = sorter(SAMPLE);
  assert.ok(steps[0] instanceof SortStep);
  assert.deepEqual(steps[steps.length - 1].array, EXPECTED);
}

test('selection sort', () => assertSorts(selectionSort));
test('bubble sort', () => assertSorts(bubbleSort));
test('insertion sort', () => assertSorts(insertionSort));
test('merge sort', () => assertSorts(mergeSort));
test('quick sort', () => assertSorts(quickSort));
test('shell sort', () => assertSorts(shellSort));
test('cocktail sort', () => assertSorts(cocktailSort));
test('heap sort', () => assertSorts(heapSort));

test('sort dispatch', () => {
  for (const [name, sorter] of Object.entries(ALGORITHMS)) {
    assert.deepEqual(sorter(SAMPLE).at(-1).array, EXPECTED, name);
  }
});

test('sort unknown algorithm', () => {
  assert.throws(() => sort(SAMPLE, 'tim'), /unknown sorting algorithm/);
});

test('input not mutated', () => {
  const values = [...SAMPLE];
  bubbleSort(values);
  assert.deepEqual(values, SAMPLE);
});

test('steps are independent copies', () => {
  const steps = bubbleSort(SAMPLE);
  steps[0].array[0] = 99;
  assert.notEqual(steps.at(-1).array[0], 99);
});
