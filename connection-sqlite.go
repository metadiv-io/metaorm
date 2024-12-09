package metaorm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite returns a new database connection for SQLite.
func SQLite(path string, silent ...bool) (*DB, error) {
	gormDB, err := gorm.Open(sqlite.Open(path), setupGormConfig(silent...))
	if err != nil {
		return nil, err
	}
	return NewDB(gormDB), nil
}

// SqliteMemory returns a new database connection for SQLite in-memory.
func SqliteMemory() (*DB, error) {
	return SQLite("file::memory:?cache=shared&busy_timeout=5000")
}
