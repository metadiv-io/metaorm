package metaorm

import "gorm.io/gorm"

func NewDB(db *gorm.DB) *DB {
	return &DB{GORM: db}
}

type DB struct {
	GORM *gorm.DB
}

// AutoMigrate runs the auto migration for the given models.
// It returns an error if the migration fails.
func (db *DB) AutoMigrate(models ...any) error {
	return db.GORM.AutoMigrate(models...)
}

// Begin starts a new transaction.
func (db *DB) Begin() *DB {
	return NewDB(db.GORM.Begin())
}

// Rollback rolls back the transaction.
func (db *DB) Rollback() *DB {
	return NewDB(db.GORM.Rollback())
}

// Commit commits the transaction.
func (db *DB) Commit() *DB {
	return NewDB(db.GORM.Commit())
}

// Preload preloads the given associations.
func (db *DB) Preload(associations ...string) *DB {
	gormDB := db.GORM
	for _, field := range associations {
		gormDB = gormDB.Preload(field)
	}
	return NewDB(gormDB)
}

// Joins joins the given associations.
func (db *DB) Joins(associations ...string) *DB {
	gormDB := db.GORM
	for _, field := range associations {
		gormDB = gormDB.Joins(field)
	}
	return NewDB(gormDB)
}

// Omit omits the given fields.
func (db *DB) Omit(fields ...string) *DB {
	return NewDB(db.GORM.Omit(fields...))
}

// Unscoped returns a new DB instance with the unscoped mode.
func (db *DB) Unscoped() *DB {
	return NewDB(db.GORM.Unscoped())
}
