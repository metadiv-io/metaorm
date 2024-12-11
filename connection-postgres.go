package metaorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres returns a new database connection for Postgres.
func Postgres(host, port, username, password, databaseName string, silent ...bool) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
	if databaseName != "" {
		dsn += fmt.Sprintf(" dbname=%s", databaseName)
	}
	gormDB, err := gorm.Open(
		postgres.Open(dsn),
		setupGormConfig(silent...),
	)
	if err != nil {
		return nil, err
	}
	return NewDB(gormDB), nil
}

// DefaultPostgres returns a new database connection for Postgres using the environment variables.
func DefaultPostgres() (*DB, error) {
	return Postgres(METAORM_HOST.String(), METAORM_PORT.String(), METAORM_USERNAME.String(), METAORM_PASSWORD.String(), METAORM_DATABASE.String())
}
