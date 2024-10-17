package demo

type tableName struct {
	tableName string
}

type InsertBuilder struct {
	tableName
	columns []string
	values []any
}
