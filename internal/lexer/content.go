package lexer

import (
	"math"
	"strings"

	"github.com/FourLineCode/minicel/internal/table"
)

func parseRows(lines []string) table.TableSlice {
	table := table.TableSlice{}

	for _, line := range lines {
		row := []string{}
		values := strings.Split(line, ",")

		for _, val := range values {
			cell := strings.Trim(val, " ")
			if cell != "" {
				row = append(row, cell)
			}
		}

		if len(row) > 0 {
			table = append(table, row)
		}
	}

	return table
}

func estimateTableSize(t table.TableSlice) table.TableSize {
	var cols int

	for _, row := range t {
		cols = int(math.Max(float64(cols), float64(len(row))))
	}

	return table.TableSize{Rows: len(t), Cols: cols}
}
