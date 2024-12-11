package metaorm

import "github.com/metadiv-io/metaorm/query"

// Query consumes the query and returns a new DB instance with the where clause applied.
func (d *DB) Query(query *query.Query) *DB {
	if query == nil {
		return d
	}
	stm, values := query.Build()
	gormDB := d.GORM
	gormDB = gormDB.Where(stm, values...)
	return NewDB(gormDB)
}

// Eq returns a new query with the equal operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) Eq(field string, value any) *query.Query {
	return query.Eq(field, value)
}

// Neq returns a new query with the not equal operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) Neq(field string, value any) *query.Query {
	return query.Neq(field, value)
}

// Gt returns a new query with the greater than operator for the given field and value.
func (db *DB) Gt(field string, value any) *query.Query {
	return query.Gt(field, value)
}

// Gte returns a new query with the greater than or equal operator for the given field and value.
func (db *DB) Gte(field string, value any) *query.Query {
	return query.Gte(field, value)
}

// Lt returns a new query with the less than operator for the given field and value.
func (db *DB) Lt(field string, value any) *query.Query {
	return query.Lt(field, value)
}

// Lte returns a new query with the less than or equal operator for the given field and value.
func (db *DB) Lte(field string, value any) *query.Query {
	return query.Lte(field, value)
}

// Like returns a new query with the like operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) Like(field string, value any) *query.Query {
	return query.Like(field, value)
}

// NotLike returns a new query with the not like operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) NotLike(field string, value any) *query.Query {
	return query.NotLike(field, value)
}

// Similar returns a new query with the similar operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) Similar(field string, value any) *query.Query {
	return query.Similar(field, value)
}

// NotSimilar returns a new query with the not similar operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) NotSimilar(field string, value any) *query.Query {
	return query.NotSimilar(field, value)
}

// In returns a new query with the in operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) In(field string, value any) *query.Query {
	return query.In(field, value)
}

// NotIn returns a new query with the not in operator for the given field and value.
// If the field is prefixed with a *, it will be decrypted.
func (db *DB) NotIn(field string, value ...any) *query.Query {
	return query.NotIn(field, value...)
}

// IsNull returns a new query with the is null operator for the given field.
func (db *DB) IsNull(field string) *query.Query {
	return query.IsNull(field)
}

// IsNotNull returns a new query with the is not null operator for the given field.
func (db *DB) IsNotNull(field string) *query.Query {
	return query.IsNotNull(field)
}

// And returns a new query with the and operator for the given queries.
func (db *DB) And(queries ...*query.Query) *query.Query {
	return query.And(queries...)
}

// Or returns a new query with the or operator for the given queries.
func (db *DB) Or(queries ...*query.Query) *query.Query {
	return query.Or(queries...)
}
