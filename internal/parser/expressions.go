package parser

import "github.com/FourLineCode/minicel/internal/table"

type Expr int
type ParsedFields map[table.CellName]ParsedFieldCell

const (
	EXPR_ADD = iota
	EXPR_SUBTRACT
	EXPR_MULTIPLY
	EXPR_DIVIDE
	EXPR_NONE
)

type Expression struct {
	Type     Expr
	Operator string
	Original table.TokenizedCell
}

type ParsedCell struct {
	Expression  Expression
	ParsedValue string
}

type ParsedFieldCell struct {
	Pos   table.Coord
	Field ParsedCell
}

type ParsedTable struct {
	Size            table.TableSize
	UnparsedSlice   table.TableSlice
	TokenizedFields table.TableFields
	ParsedFields    ParsedFields
}

func New(t table.TokenizedTable, fields ParsedFields) ParsedTable {
	return ParsedTable{
		Size:            t.Size,
		UnparsedSlice:   t.Slice,
		TokenizedFields: t.Fields,
		ParsedFields:    fields,
	}
}
