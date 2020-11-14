package main

import (
	"io"
)

func insertionSort(array []int, writer io.Writer, reverse bool) []int {
	// display original data.
	if writer != nil {
		visualize(writer, array)
	}

	for index, num := range array {
		prevIndex := index - 1

		for prevIndex >= 0 && isMustSwap(reverse, array[prevIndex], num) {
			array[prevIndex+1] = array[prevIndex]
			array[prevIndex] = num

			prevIndex--
			// only display bar chart when writter has setted.
			if writer != nil {
				visualize(writer, array)
			}
		}
	}

	return array
}

// isMustSwap determines that if current position must be swapped with previous position.
func isMustSwap(isReverseMode bool, prev, current int) bool {
	if isReverseMode {
		return prev < current
	}

	return prev > current
}

func same(array1, array2 []int) bool {
	// compared array must having same length.
	if len(array1) != len(array2) {
		return false
	}

	for index, num := range array1 {
		if num != array2[index] {
			return false
		}
	}

	return true
}
