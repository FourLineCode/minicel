package parser

import (
	"fmt"
	"math"
	"os"

	"github.com/FourLineCode/minicel/internal/table"
	"github.com/FourLineCode/minicel/internal/utils"
)

func (t ParsedTable) PrintTable() {
	columnLengths := make([]int, t.Size.Cols)
	sum := 0
	minSpaces := 3
	colSeparator := " | "

	for i := 0; i < t.Size.Cols; i++ {
		columnLengths[i] = t.maxColumnLength(i)
		sum += int(math.Max(float64(minSpaces), float64(columnLengths[i])))
	}

	extraLength := int(math.Floor(math.Sqrt(float64(len(colSeparator)))))
	tableLength := sum + t.Size.Cols*len(colSeparator) - extraLength

	fmt.Println("Parsed Table:")
	printSeparators(tableLength)

	for rowIndex, row := range t.UnparsedSlice {
		for i := 0; i < t.Size.Cols; i++ {
			length := columnLengths[i]
			spaces := int(math.Max(float64(minSpaces), float64(length)))

			if i < len(row) {
				cellName := utils.GetCellnameFromIndex(rowIndex, i)
				cell := t.ParsedFields[table.CellName(cellName)].Field.ParsedValue
				fmt.Printf("%v", cell)
				spaces -= len(cell)
			}
			for i := 0; i < spaces; i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("%v", colSeparator)
		}
		fmt.Println()
	}

	printSeparators(tableLength)
}

func (t ParsedTable) maxColumnLength(index int) int {
	if index < 0 || index >= t.Size.Cols {
		fmt.Println("Error: column index out of bound", index)
		os.Exit(1)
	}

	max := math.Inf(-1)
	for rowIndex, row := range t.UnparsedSlice {
		if index >= len(row) {
			continue
		}
		cellName := utils.GetCellnameFromIndex(rowIndex, index)
		max = math.Max(max, float64(len(t.ParsedFields[table.CellName(cellName)].Field.ParsedValue)))
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
