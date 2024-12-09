package metaorm

import (
	"fmt"

	"github.com/metadiv-io/metaorm/internal/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL returns a new database connection for MySQL.
func MySQL(host, port, username, password, databaseName string, silent ...bool) (*database.DB, error) {
	gormDB, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				username, password, host, port, databaseName,
			),
		),
		setupGormConfig(silent...),
	)
	if err != nil {
		return nil, err
	}
	return database.New(gormDB), nil
}

// DefaultMySQL returns a new database connection for MySQL using the environment variables.
func DefaultMySQL() (*database.DB, error) {
	return MySQL(METAORM_HOST.String(), METAORM_PORT.String(), METAORM_USERNAME.String(), METAORM_PASSWORD.String(), METAORM_DATABASE.String())
}