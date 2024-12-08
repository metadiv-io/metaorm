package builder

import (
	"github.com/metadiv-io/metaorm/internal/operator"
	"github.com/metadiv-io/metaorm/internal/query"
)

// QueryBuilder is a builder for building queries.
type QueryBuilder struct {
	Operator operator.Operator
	Queries  []*query.Query
}

// Add adds a query to the builder.
func (qb *QueryBuilder) Add(query *query.Query) *QueryBuilder {
	if query == nil {
		return qb
	}
	qb.Queries = append(qb.Queries, query)
	return qb
}

// Build builds the query.
func (qb *QueryBuilder) Build() *query.Query {
	if len(qb.Queries) == 0 {
		return nil
	}
	if qb.Operator.IsAnd() {
		return query.And(qb.Queries...)
	}
	return query.Or(qb.Queries...)
}
