package table

import (
	"fmt"
	"math"
	"os"
)

func (t Table) PrintSize() {
	fmt.Printf("Table Size: %vx%v\n", t.Size.Rows, t.Size.Cols)
}

func (t Table) PrintSlice() {
	columnLengths := make([]int, t.Size.Cols)
	sum := 0

	for i := 0; i < t.Size.Cols; i++ {
		columnLengths[i] = t.maxColumnLength(i)
		sum += columnLengths[i]
	}

	tableLength := sum + t.Size.Cols*3 - 1

	printSeparators(tableLength)

	for _, row := range t.Slice {
		for i := 0; i < t.Size.Cols; i++ {
			length := columnLengths[i]
			spaces := length

			if i < len(row) {
				cell := row[i]
				fmt.Printf("%v", cell)
				spaces = length - len(cell)
			}
			for i := 0; i < spaces; i++ {
				fmt.Printf(" ")
			}
			fmt.Printf(" | ")
		}
		fmt.Println()
	}

	printSeparators(tableLength)
}

func (t Table) maxColumnLength(index int) int {
	if index < 0 || index >= t.Size.Cols {
		fmt.Println("Error: column index out of bound", index)
		os.Exit(1)
	}

	max := math.Inf(-1)
	for _, row := range t.Slice {
		if index >= len(row) {
			continue
		}
		max = math.Max(max, float64(len(row[index])))
	}

	if max < 0 {
		fmt.Println("Error: error while calculating max column size")
		os.Exit(1)
	}

	return int(max)
}

func printSeparators(length int) {
	for i := 0; i < length; i++ {
		fmt.Printf("=")
	}
	fmt.Println()
}
