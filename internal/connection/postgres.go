package connection

import (
	"fmt"

	"github.com/metadiv-io/metaorm/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres(host, port, username, password, databaseName string, silent ...bool) (*database.DB, error) {
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
	return database.New(gormDB), nil
}
