package metaorm

import (
	"github.com/metadiv-io/metaorm/internal/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite returns a new database connection for SQLite.
func SQLite(path string, silent ...bool) (*database.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open(path), setupGormConfig(silent...))
	if err != nil {
		return nil, err
	}
	return database.New(gormDB), nil
}

// SqliteMemory returns a new database connection for SQLite in-memory.
func SqliteMemory() (*database.DB, error) {
	return SQLite("file::memory:?cache=shared&busy_timeout=5000")
}
