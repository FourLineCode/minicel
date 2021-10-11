package lexer

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/FourLineCode/minicel/internal/table"
)

func parseRows(lines []string) table.TableSlice {
	table := table.TableSlice{}

	for index, line := range lines {
		if strings.Contains(line, ";") && strings.Contains(line, ",") {
			fmt.Println("Error: invalid syntaxt for .csv file at line", index)
			fmt.Println("\t.csv files cannot contain both comma (,) and semicolon (;)")
			os.Exit(1)
		}

		delim := ","
		if !strings.Contains(line, ",") && strings.Contains(line, ";") {
			delim = ";"
		}

		row := []string{}
		values := strings.Split(line, delim)

		for _, val := range values {
			cell := strings.TrimSpace(val)
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
