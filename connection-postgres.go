package metaorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres returns a new database connection for Postgres.
func Postgres(host, port, username, password, databaseName string, silent ...bool) (*DB, error) {
	gormDB, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				host, port, username, databaseName, password,
			),
		),
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
