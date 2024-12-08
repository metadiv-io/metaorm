package query

import "github.com/metadiv-io/metaorm/internal/operator"

func Eq(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Eq(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedEq(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Eq(),
		Value:     value,
		Encrypted: true,
	}
}

func Neq(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Neq(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedNeq(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Neq(),
		Value:     value,
		Encrypted: true,
	}
}

func Gt(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Gt(),
		Value:     value,
		Encrypted: false,
	}
}

func Gte(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Gte(),
		Value:     value,
		Encrypted: false,
	}
}

func Lt(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Lt(),
		Value:     value,
		Encrypted: false,
	}
}

func Lte(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Lte(),
		Value:     value,
		Encrypted: false,
	}
}

func Like(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Like(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedLike(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Like(),
		Value:     value,
		Encrypted: true,
	}
}

func NotLike(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.NotLike(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedNotLike(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.NotLike(),
		Value:     value,
		Encrypted: true,
	}
}

func Similar(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Similar(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedSimilar(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.Similar(),
		Value:     value,
		Encrypted: true,
	}
}

func NotSimilar(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.NotSimilar(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedNotSimilar(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.NotSimilar(),
		Value:     value,
		Encrypted: true,
	}
}

func In(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.In(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedIn(field string, value any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.In(),
		Value:     value,
		Encrypted: true,
	}
}

func NotIn(field string, value ...any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.NotIn(),
		Value:     value,
		Encrypted: false,
	}
}

func DecryptedNotIn(field string, value ...any) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.NotIn(),
		Value:     value,
		Encrypted: true,
	}
}

func IsNull(field string) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.IsNull(),
		Value:     nil,
		Encrypted: false,
	}
}

func IsNotNull(field string) *Query {
	return &Query{
		Field:     SafeField(field),
		Operator:  operator.IsNotNull(),
		Value:     nil,
		Encrypted: false,
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
