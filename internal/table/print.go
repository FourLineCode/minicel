package table

import (
	"fmt"
	"math"
	"os"

	"github.com/FourLineCode/minicel/internal/utils"
)

func (t TokenizedTable) PrintSize() {
	fmt.Printf("Table Size: %vx%v\n", t.Size.Rows, t.Size.Cols)
}

func (t TokenizedTable) PrintSlice() {
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

	fmt.Println("Table Slice:")
	printSeparators(tableLength)

	for _, row := range t.Slice {
		for i := 0; i < t.Size.Cols; i++ {
			length := columnLengths[i]
			spaces := int(math.Max(float64(minSpaces), float64(length)))

			if i < len(row) {
				cell := row[i]
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

func (t TokenizedTable) PrintTokens() {
	tokenToTypename := map[int]string{
		0: "TEXT",
		1: "NUMBER",
		2: "EXPR",
	}

	columnLengths := make([]int, t.Size.Cols)
	sum := 0
	minSpaces := 0
	colSeparator := " | "

	for _, v := range tokenToTypename {
		minSpaces = int(math.Max(float64(minSpaces), float64(len(v))))
	}

	for i := 0; i < t.Size.Cols; i++ {
		columnLengths[i] = t.maxColumnLength(i)
		sum += int(math.Max(float64(minSpaces), float64(columnLengths[i])))
	}

	extraLength := int(math.Floor(math.Sqrt(float64(len(colSeparator)))))
	tableLength := sum + t.Size.Cols*len(colSeparator) - extraLength

	fmt.Println("Table Tokens:")
	printSeparators(tableLength)

	for rowIndex, row := range t.Slice {
		for i := 0; i < t.Size.Cols; i++ {
			length := columnLengths[i]
			spaces := int(math.Max(float64(minSpaces), float64(length)))

			cellName := utils.GetCellnameFromIndex(rowIndex, i)
			cell := t.Fields[CellName(cellName)]

			cellType := tokenToTypename[int(cell.Field.Type)]
			if i < len(row) {
				fmt.Printf("%v", cellType)
				spaces -= len(cellType)
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

func (t TokenizedTable) maxColumnLength(index int) int {
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
