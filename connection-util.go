package metaorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupGormConfig(silent ...bool) (config *gorm.Config) {
	config = &gorm.Config{}
	if len(silent) > 0 && silent[0] {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return
}
