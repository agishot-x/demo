package demo

type tableName struct {
	tableName string
}

type insertBuilder struct {
	tableName
	columns []string
	values []any
}