package main

import (
	"bytes"
	"io"
	"strconv"
)

// visualize display given slices into vertical bar chart (with label)
func visualize(writer io.Writer, list []int) {
	highest := maxValue(list)

	var buf bytes.Buffer

	// writes chart bar.
	for ; highest > 0; highest-- {
		for _, num := range list {
			if num >= highest {
				buf.WriteString("|")
			} else {
				buf.WriteString(" ")
			}

			buf.WriteString(" ") // used as right margin.
		}

		buf.WriteString("\n")
	}

	// writes chart label.
	for _, num := range list {
		buf.WriteString(strconv.Itoa(num))
		buf.WriteString(" ")
	}

	buf.WriteString("\n\n")
	buf.WriteTo(writer)
}

// maxValue returns highest / maximum element's value from given slices.
func maxValue(list []int) int {
	var highest int
	for _, val := range list {
		if val > highest {
			highest = val
		}
	}

	return highest
}
