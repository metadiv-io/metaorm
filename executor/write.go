package executor

import (
	"github.com/metadiv-io/metaorm"
	"github.com/metadiv-io/metaorm/query"
)

func Save[Model any](db *metaorm.DB, model *Model) (*Model, error) {
	if err := db.GORM.Save(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func SaveAll[Model any](db *metaorm.DB, models []Model) ([]Model, error) {
	if err := db.GORM.Save(models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func Delete[Model any](db *metaorm.DB, model *Model) error {
	if err := db.GORM.Delete(model).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAll[Model any](db *metaorm.DB, models []Model) error {
	tx := db.Begin()
	for i := range models {
		if err := tx.GORM.Delete(&models[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.GORM.Commit().Error
}

func DeleteByQuery[Model any](db *metaorm.DB, query *query.Query) error {
	db = db.Query(query)
	return db.GORM.Delete(new(Model)).Error
}
