package metaorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL returns a new database connection for MySQL.
func MySQL(host, port, username, password, databaseName string, silent ...bool) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)", username, password, host, port)
	if databaseName != "" {
		dsn += fmt.Sprintf("/%s", databaseName)
	} else {
		dsn += "/"
	}
	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(
		mysql.Open(dsn),
		setupGormConfig(silent...),
	)
	if err != nil {
		return nil, err
	}
	return NewDB(gormDB), nil
}

// DefaultMySQL returns a new database connection for MySQL using the environment variables.
func DefaultMySQL() (*DB, error) {
	return MySQL(METAORM_HOST.String(), METAORM_PORT.String(), METAORM_USERNAME.String(), METAORM_PASSWORD.String(), METAORM_DATABASE.String())
}
