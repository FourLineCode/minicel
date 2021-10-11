package parser

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FourLineCode/minicel/internal/table"
)

func Parse(tokenizedTable table.TokenizedTable) ParsedTable {
	parsedTable := ParsedTable{
		Size:            tokenizedTable.Size,
		UnparsedSlice:   tokenizedTable.Slice,
		TokenizedFields: tokenizedTable.Fields,
		ParsedFields:    ParsedFields{},
	}

	for key, val := range tokenizedTable.Fields {
		parsedFieldCell := ParsedFieldCell{Pos: val.Pos}

		expression := Expression{Original: val.Field}
		if val.Field.Type != table.TOKEN_EXPR {
			expression.Type = EXPR_NONE
			parsedFieldCell.Field = ParsedCell{Expression: expression, ParsedValue: val.Field.Value}
			parsedTable.ParsedFields[key] = parsedFieldCell
		}
	}

	for key, val := range tokenizedTable.Fields {
		if val.Field.Type != table.TOKEN_EXPR {
			continue
		}

		parsedFieldCell := ParsedFieldCell{Pos: val.Pos}
		expression := Expression{Original: val.Field}

		var err error
		expression.Type, expression.Operator, err = parseExressionType(val.Field.Value)
		if err != nil {
			fmt.Println("Error: invalid expression type on field", key)
			validExpressions()
			os.Exit(1)
		}

		parsedValue, err := parseExpressionValue(expression, parsedTable)
		if err != nil {
			fmt.Println("Error: expression has invalid operation on field", key)
			validExpressions()
			os.Exit(1)
		}

		parsedFieldCell.Field = ParsedCell{Expression: expression, ParsedValue: parsedValue}

		parsedTable.ParsedFields[key] = parsedFieldCell
	}

	return parsedTable
}

func parseExressionType(val string) (Expr, string, error) {
	valid := isValidExpression(val)
	if !valid {
		return EXPR_NONE, "", errors.New("invalid expression")
	}

	if strings.Contains(val, "+") && len(strings.Split(val, "+")) == 2 {
		return EXPR_ADD, "+", nil
	} else if strings.Contains(val, "-") && len(strings.Split(val, "-")) == 2 {
		return EXPR_SUBTRACT, "-", nil
	} else if strings.Contains(val, "*") && len(strings.Split(val, "*")) == 2 {
		return EXPR_MULTIPLY, "*", nil
	} else if strings.Contains(val, "/") && len(strings.Split(val, "/")) == 2 {
		return EXPR_DIVIDE, "/", nil
	}

	return EXPR_NONE, "", nil
}

func validExpressions() {
	fmt.Println("Valid Expressions:")
	fmt.Println("\tA1+B1 - Add")
	fmt.Println("\tA1-B1 - Subtract")
	fmt.Println("\tA1*B1 - Multiply")
	fmt.Println("\tA1/B1 - Divide")
}

func isValidExpression(val string) bool {
	plus := strings.Contains(val, "+")
	minus := strings.Contains(val, "-")
	mul := strings.Contains(val, "*")
	div := strings.Contains(val, "/")

	return plus || minus || mul || div
}

func parseExpressionValue(exp Expression, t ParsedTable) (string, error) {
	cells := strings.Split(exp.Original.Value, exp.Operator)
	cell1Name, cell2Name := cells[0], cells[1]

	field1 := t.ParsedFields[table.CellName(cell1Name)].Field
	field2 := t.ParsedFields[table.CellName(cell2Name)].Field

	// TODO: lets assume expression mention only numbers
	if field1.Expression.Original.Type != table.TOKEN_NUMBER || field2.Expression.Original.Type != table.TOKEN_NUMBER {
		fmt.Println("Error: expressions must point to fields of NUMBER type")
		os.Exit(1)
	}

	cell1, cell2 := field1.ParsedValue, field2.ParsedValue

	num1, err := strconv.ParseFloat(cell1, 64)
	if err != nil {
		fmt.Println("Error: invalid number in expression")
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(cell2, 64)
	if err != nil {
		fmt.Println("Error: invalid number in expression")
		os.Exit(1)
	}

	var num float64
	switch exp.Type {
	case EXPR_ADD:
		num = num1 + num2
	case EXPR_SUBTRACT:
		num = num1 - num2
	case EXPR_MULTIPLY:
		num = num1 * num2
	case EXPR_DIVIDE:
		num = num1 / num2
	case EXPR_NONE:
		fmt.Println("Error: unreachable code (EXPR_NONE)")
		os.Exit(1)
	default:
		fmt.Println("Error: unreachable code (EXPR_NONE)")
		os.Exit(1)
	}

	value := fmt.Sprintf("%f", num)
	fmt.Println(value)
	return value, nil
}
