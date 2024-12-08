package metaorm

import (
	"github.com/metadiv-io/metaorm/internal/database"
	"github.com/metadiv-io/metaorm/internal/query"
)

// Page is a pagination model.
type Page struct {
	Page  int `json:"page" form:"page"`
	Size  int `json:"size" form:"size"`
	Total int `json:"total" form:"-"`
}

// Consume consumes the page and returns a new database.DB instance with the offset and limit applied.
func (p *Page) Consume(db *database.DB) *database.DB {
	gormDB := db.GORM
	if p.Page > 0 {
		gormDB = gormDB.Offset((p.Page - 1) * p.Size)
	}
	if p.Size > 0 {
		gormDB = gormDB.Limit(p.Size)
	}
	return database.New(gormDB)
}

// Sort is a sort model.
type Sort struct {
	Field string `json:"field" form:"field"`
	Asc   bool   `json:"asc" form:"asc"`
}

// Consume consumes the sort and returns a new database.DB instance with the order applied.
func (s *Sort) Consume(db *database.DB) *database.DB {
	gormDB := db.GORM
	if s.Field != "" {
		if s.Asc {
			gormDB = gormDB.Order(query.SafeField(s.Field) + " ASC")
		} else {
			gormDB = gormDB.Order(query.SafeField(s.Field) + " DESC")
		}
	}
	return database.New(gormDB)
}
