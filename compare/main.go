package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	source, target := args[0], args[1]
	if source == "" || target == "" {
		log.Fatal("source and target are required")
	}

	diff, err := Compare(source, target)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range diff {
		fmt.Println(val)
	}
}
