package ui


type Table struct {
	Type  string
	Width []int
	Rows  []TableRow
}

func NewTable() Table {
	return Table{Type: "table"}
}

type TableRow struct {
	Type  string
	Color string
	Row   []String
}

func NewTableRow() TableRow {
	return TableRow{Type: "table_row"}
}
