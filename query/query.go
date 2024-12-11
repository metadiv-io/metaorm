package query

import (
	"fmt"
	"os"
	"strings"

	"github.com/metadiv-io/metaorm/internal/operator"
)

type Query struct {
	Field     string
	Operator  operator.Operator
	Value     any
	Encrypted bool
	Children  []*Query
}

func (q *Query) Build() (clause string, values []any) {
	return q.build(q, make([]any, 0))
}

func (q *Query) build(query *Query, values []any) (string, []any) {
	var stm string
	if query.Operator.IsAnd() || query.Operator.IsOr() {
		return q.buildWithChildren(query, values)
	} else if query.Operator.IsNull() || query.Operator.IsNotNull() {
		stm = query.Field + " " + query.Operator.Operator
	} else if query.Operator.IsIn() || query.Operator.IsNotIn() {
		stm = q.getField(query.Field, query.Encrypted) + " " + query.Operator.Operator + " (" + strings.TrimRight(strings.Repeat("?,", len(query.Value.([]any))), ",") + ")"
		values = append(values, query.Value.([]any)...)
	} else if query.Operator.IsSimilar() || query.Operator.IsNotSimilar() {
		stm = q.getField(query.Field, query.Encrypted) + " " + operator.Like().Operator + " ?"
		values = append(values, fmt.Sprintf("%%%s%%", query.Value))
	} else {
		stm = q.getField(query.Field, query.Encrypted) + " " + query.Operator.Operator + " ?"
		values = append(values, query.Value)
	}
	return stm, values
}

func (q *Query) getField(field string, encrypted bool) string {
	if encrypted {
		return "AES_DECRYPT(" + field + ", '" + os.Getenv("METAORM_ENCRYPT_KEY") + "')"
	}
	return field
}

func (q *Query) buildWithChildren(query *Query, values []any) (string, []any) {
	var buf []string
	for _, child := range query.Children {
		childStmt, childValues := q.build(child, values)
		buf = append(buf, childStmt)
		values = childValues
	}
	stm := "(" + strings.Join(buf, " "+query.Operator.Operator+" ") + ")"
	return stm, values
}
