package demo

type Column struct {
	name string
	value any
}

type ColumnOp struct {
	Column
	op CompareType
}

type builder struct {
	tableName string
	columns []Column
}

func (b builder) Size() int {
	return len(b.columns)
}

type whereBuilder struct {
	where []ColumnOp
}

type InsertBuilder struct {
	builder
}

type UpdateBuilder struct {
	builder
	whereBuilder
}

type DeleteBuilder struct {
	tableName string
	whereBuilder
}

type QueryBuilder struct {
	builder
	whereBuilder
}
