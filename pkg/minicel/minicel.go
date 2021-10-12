package minicel

import (
	"io/ioutil"

	"github.com/FourLineCode/minicel/internal/lexer"
	"github.com/FourLineCode/minicel/internal/parser"
)

func ParseCSV(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	tokenizedTable := lexer.Parse(string(content))

	tokenizedTable.PrintSize()
	tokenizedTable.PrintSlice()
	// tokenizedTable.PrintTokens()

	parsedTable := parser.Parse(tokenizedTable)
	parsedTable.PrintTable()

	return nil
}
