package utils

import "strconv"

func GetColnameFromIndex(row, col int) string {
	char := string(rune((col % 26) + 65))
	iter := col / 26
	colName := char
	for i := 0; i < iter; i++ {
		colName += char
	}

	return colName
}

func GetCellnameFromIndex(row, col int) string {
	colName := GetColnameFromIndex(row, col)

	return colName + strconv.Itoa(row)
}
