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
