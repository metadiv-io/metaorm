package metaorm

import (
	"github.com/metadiv-io/metaorm/internal/builder"
	"github.com/metadiv-io/metaorm/internal/operator"
	"github.com/metadiv-io/metaorm/internal/query"
)

// NewAndQueryBuilder returns a new and query builder.
// The default operator is and.
func NewAndQueryBuilder() *builder.QueryBuilder {
	return &builder.QueryBuilder{
		Operator: operator.And(),
		Queries:  make([]*query.Query, 0),
	}
}

// NewOrQueryBuilder returns a new or query builder.
// The default operator is or.
func NewOrQueryBuilder() *builder.QueryBuilder {
	return &builder.QueryBuilder{
		Operator: operator.Or(),
		Queries:  make([]*query.Query, 0),
	}
}
