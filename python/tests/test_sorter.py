import unittest

from sortviz import (
    ALGORITHMS,
    SortStep,
    bubble_sort,
    cocktail_sort,
    heap_sort,
    insertion_sort,
    merge_sort,
    quick_sort,
    selection_sort,
    shell_sort,
    sort,
)

SAMPLE = [5, 3, 8, 4, 2]
EXPECTED = [2, 3, 4, 5, 8]


class SorterTest(unittest.TestCase):
    def assert_sorts(self, sorter) -> None:
        steps = sorter(SAMPLE)
        self.assertIsInstance(steps[0], SortStep)
        self.assertEqual(steps[-1].array, EXPECTED)

    def test_selection_sort(self) -> None:
        self.assert_sorts(selection_sort)

    def test_bubble_sort(self) -> None:
        self.assert_sorts(bubble_sort)

    def test_insertion_sort(self) -> None:
        self.assert_sorts(insertion_sort)

    def test_merge_sort(self) -> None:
        self.assert_sorts(merge_sort)

    def test_quick_sort(self) -> None:
        self.assert_sorts(quick_sort)

    def test_shell_sort(self) -> None:
        self.assert_sorts(shell_sort)

    def test_cocktail_sort(self) -> None:
        self.assert_sorts(cocktail_sort)

    def test_heap_sort(self) -> None:
        self.assert_sorts(heap_sort)

    def test_sort_dispatch(self) -> None:
        for name, sorter in ALGORITHMS.items():
            with self.subTest(algorithm=name):
                self.assertEqual(sort(SAMPLE, name)[-1].array, EXPECTED)

    def test_sort_unknown_algorithm(self) -> None:
        with self.assertRaises(ValueError):
            sort(SAMPLE, "tim")

    def test_input_not_mutated(self) -> None:
        values = list(SAMPLE)
        bubble_sort(values)
        self.assertEqual(values, SAMPLE)

    def test_steps_are_independent_copies(self) -> None:
        steps = bubble_sort(SAMPLE)
        steps[0].array[0] = 99
        self.assertNotEqual(steps[-1].array[0], 99)


if __name__ == "__main__":
    unittest.main()
