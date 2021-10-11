package minicel

import (
	"io/ioutil"

	"github.com/FourLineCode/minicel/internal/lexer"
)

func ParseCSV(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	t := lexer.Parse(string(content))

	t.PrintSize()
	t.PrintSlice()
	t.PrintTokens()

	return nil
}
