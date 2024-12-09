package executor

import (
	"github.com/metadiv-io/metaorm"
	"github.com/metadiv-io/metaorm/query"
)

func FindOne[Model any](db *metaorm.DB, query *query.Query) (*Model, error) {
	db = db.Query(query)
	output := new(Model)
	if err := db.GORM.First(output).Error; err != nil {
		return nil, err
	}
	return output, nil
}

func FindAll[Model any](db *metaorm.DB, query *query.Query) ([]Model, error) {
	db = db.Query(query)
	var output []Model
	if err := db.GORM.Find(&output).Error; err != nil {
		return nil, err
	}
	return output, nil
}

func Count[Model any](db *metaorm.DB, query *query.Query) (int64, error) {
	db = db.Query(query)
	var total int64
	if err := db.GORM.Model(new(Model)).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func FindAllComplex[Model any](db *metaorm.DB, query *query.Query, page *metaorm.Page, sort *metaorm.Sort) ([]Model, *metaorm.Page, error) {
	// Part 1: Find Records
	tx1 := db.Begin()
	tx1 = tx1.Query(query)
	tx1 = page.Consume(tx1)
	tx1 = sort.Consume(tx1)
	output := make([]Model, 0)
	if err := tx1.GORM.Find(&output).Error; err != nil {
		tx1.Rollback()
		return nil, nil, err
	}
	tx1.Commit()

	// Part 2: Count Records
	if page != nil {
		tx2 := db.Begin()
		tx2 = tx2.Query(query)
		var total int64
		if err := tx2.GORM.Model(&output).Count(&total).Error; err != nil {
			tx2.Rollback()
			return nil, nil, err
		}
		page.Total = int(total)
		tx2.Commit()
	}

	return output, page, nil
}
