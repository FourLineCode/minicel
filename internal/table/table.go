package table

type CellName string
type TableSlice [][]string
type TableFields map[CellName]TableCell

type Coord struct {
	Row int
	Col string
}

type TokenizedCell struct {
	Type  Token
	Value string
}

type TableCell struct {
	Pos   Coord
	Field TokenizedCell
}

type TableSize struct {
	Rows int
	Cols int
}

type TokenizedTable struct {
	Size   TableSize
	Slice  TableSlice
	Fields TableFields
}

func New(size TableSize, slice TableSlice, fields TableFields) TokenizedTable {
	return TokenizedTable{
		Size:   size,
		Slice:  slice,
		Fields: fields,
	}
}
