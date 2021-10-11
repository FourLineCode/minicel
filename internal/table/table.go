package table

type Token int
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

type Table struct {
	Size   TableSize
	Slice  TableSlice
	Fields TableFields
}

func New(size TableSize, slice TableSlice, fields TableFields) Table {
	return Table{
		Size:   size,
		Slice:  slice,
		Fields: fields,
	}
}
