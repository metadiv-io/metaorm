package metaorm

import (
	"github.com/metadiv-io/metaorm/internal/connection"
	"github.com/metadiv-io/metaorm/internal/database"
)

// MySQL returns a new database connection for MySQL.
func MySQL(host, port, username, password, databaseName string, silent ...bool) (*database.DB, error) {
	return connection.MySQL(host, port, username, password, databaseName, silent...)
}

// DefaultMySQL returns a new database connection for MySQL using the environment variables.
func DefaultMySQL() (*database.DB, error) {
	return MySQL(METAORM_HOST.String(), METAORM_PORT.String(), METAORM_USERNAME.String(), METAORM_PASSWORD.String(), METAORM_DATABASE.String())
}

// Postgres returns a new database connection for Postgres.
func Postgres(host, port, username, password, databaseName string, silent ...bool) (*database.DB, error) {
	return connection.Postgres(host, port, username, password, databaseName, silent...)
}

// DefaultPostgres returns a new database connection for Postgres using the environment variables.
func DefaultPostgres() (*database.DB, error) {
	return Postgres(METAORM_HOST.String(), METAORM_PORT.String(), METAORM_USERNAME.String(), METAORM_PASSWORD.String(), METAORM_DATABASE.String())
}

// SQLite returns a new database connection for SQLite.
func SQLite(path string, silent ...bool) (*database.DB, error) {
	return connection.SQLite(path, silent...)
}

// SqliteMemory returns a new database connection for SQLite in-memory.
func SqliteMemory() (*database.DB, error) {
	return connection.SqliteMemory()
}
