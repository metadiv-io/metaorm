package query

import "github.com/metadiv-io/metaorm/internal/operator"

func Eq(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Eq(),
		Value:    value,
	}
}

func Neq(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Neq(),
		Value:    value,
	}
}

func Gt(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Gt(),
		Value:    value,
	}
}

func Gte(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Gte(),
		Value:    value,
	}
}

func Lt(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Lt(),
		Value:    value,
	}
}

func Lte(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Lte(),
		Value:    value,
	}
}

func Like(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Like(),
		Value:    value,
	}
}

func NotLike(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.NotLike(),
		Value:    value,
	}
}

func Similar(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.Similar(),
		Value:    value,
	}
}

func NotSimilar(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.NotSimilar(),
		Value:    value,
	}
}

func In(field string, value any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.In(),
		Value:    value,
	}
}

func NotIn(field string, value ...any) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.NotIn(),
		Value:    value,
	}
}

func IsNull(field string) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.IsNull(),
		Value:    nil,
	}
}

func IsNotNull(field string) *Query {
	return &Query{
		Field:    SafeField(field),
		Operator: operator.IsNotNull(),
		Value:    nil,
	}
}

func And(queries ...*Query) *Query {
	return &Query{
		Operator: operator.And(),
		Children: queries,
	}
}

func Or(queries ...*Query) *Query {
	return &Query{
		Operator: operator.Or(),
		Children: queries,
	}
}
