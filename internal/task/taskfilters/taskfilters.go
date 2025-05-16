package taskfilters

import "fmt"

type TaskFilter interface {
	getColumnName() string
	getColumnValue() string
	GetQueryClause() string
}

type StringMatchTaskFilter struct {
	ColumnName  string
	ColumnValue string
}

func (f *StringMatchTaskFilter) getColumnName() string {
	return f.ColumnName
}

func (f *StringMatchTaskFilter) getColumnValue() string {
	return f.ColumnValue
}

func (f *StringMatchTaskFilter) GetQueryClause() string {
	return fmt.Sprintf("%s = '%s'", f.ColumnName, f.ColumnValue)
}
