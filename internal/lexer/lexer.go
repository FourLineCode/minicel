package lexer

import (
	"strings"

	"github.com/FourLineCode/minicel/internal/table"
)

func Parse(content string) table.Table {
	lines := strings.Split(content, "\n")

	slice := parseRows(lines)
	size := estimateTableSize(slice)
	fields := tokenizeTable(slice)

	t := table.New(size, slice, fields)

	return t
}
