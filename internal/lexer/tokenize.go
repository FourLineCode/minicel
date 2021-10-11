package lexer

import (
	"strconv"
	"strings"

	"github.com/FourLineCode/minicel/internal/table"
	"github.com/FourLineCode/minicel/internal/utils"
)

func tokenizeTable(slice table.TableSlice) table.TableFields {
	fields := make(table.TableFields)

	for rowIndex, row := range slice {
		for colIndex, col := range row {
			colName := utils.GetColnameFromIndex(rowIndex, colIndex)
			cellName := utils.GetCellnameFromIndex(rowIndex, colIndex)
			cellPos := table.Coord{Row: rowIndex, Col: colName}

			tokenizedCell := table.TokenizedCell{}
			if strings.HasPrefix(col, "=") {
				tokenizedCell.Type = table.TOKEN_EXPR
				tokenizedCell.Value = strings.TrimLeft(col, "=")
			} else if isNumeric(col) {
				tokenizedCell.Type = table.TOKEN_NUMBER
				tokenizedCell.Value = col
			} else {
				tokenizedCell.Type = table.TOKEN_TEXT
				tokenizedCell.Value = col
			}

			tableCell := table.TableCell{
				Pos:   cellPos,
				Field: tokenizedCell,
			}

			fields[table.CellName(cellName)] = tableCell
		}
	}

	return fields
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
