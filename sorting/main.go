package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		output  string
		numbers string
		reverse bool
	)

	flag.StringVar(&output, "output", "", "output file. if not set, default output is os.Stdout")
	flag.StringVar(&numbers, "numbers", "", "number sets, separated with comma. e.g. 1,3,4,2,3")
	flag.BoolVar(&reverse, "reverse", false, "reverse sorting, default false")

	flag.Parse()

	list := convertInt(strings.Split(numbers, ","))

	// check if user wants to export output into a file.
	if output != "" {
		// write output to file.
		if err := outputToFile(list, output, reverse); err != nil {
			log.Fatal(err)
		}

		return
	}

	// write to terminal as default output.
	insertionSort(list, os.Stdout, reverse)
}

// outputToFile stores result into a file.
func outputToFile(list []int, filePath string, reverse bool) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	insertionSort(list, file, reverse)

	return nil
}

// convertInt converts string slices into int slices.
func convertInt(args []string) []int {
	var result []int
	for _, val := range args {
		val = strings.Trim(val, " ")
		if converted64, err := strconv.ParseInt(val, 10, 32); err == nil {
			result = append(result, int(converted64))
		}
	}

	return result
}
