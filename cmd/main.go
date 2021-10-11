package main

import (
	"fmt"
	"os"

	"github.com/FourLineCode/minicel/pkg/minicel"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: input file path not provided")
		usage()
		os.Exit(1)
	}

	filePath := os.Args[1]

	err := minicel.ParseCSV(filePath)
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Println("Usage: ./minicel <input.csv> [OPTIONS]")
}
