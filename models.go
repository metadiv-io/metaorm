package metaorm

import (
	"os"
	"strings"

	"github.com/metadiv-io/metaorm/query"
)

// Page is a pagination model.
type Page struct {
	Page  int `json:"page" form:"page"`
	Size  int `json:"size" form:"size"`
	Total int `json:"total" form:"-"`
}

// Consume consumes the page and returns a new DB instance with the offset and limit applied.
func (p *Page) Consume(db *DB) *DB {
	gormDB := db.GORM
	if p.Page > 0 {
		gormDB = gormDB.Offset((p.Page - 1) * p.Size)
	}
	if p.Size > 0 {
		gormDB = gormDB.Limit(p.Size)
	}
	return NewDB(gormDB)
}

// Sort is a sort model.
// add "*" to the field to sort by encrypted field
type Sort struct {
	Field string `json:"field" form:"field"`
	Asc   bool   `json:"asc" form:"asc"`
}

// Consume consumes the sort and returns a new DB instance with the order applied.
func (s *Sort) Consume(db *DB) *DB {
	gormDB := db.GORM
	if s.Field != "" {
		if strings.HasPrefix(s.Field, "*") {
			s.Field = strings.TrimPrefix(s.Field, "*")
			if s.Asc {
				gormDB = gormDB.Order("AES_DECRYPT(" + query.SafeField(s.Field) + ", '" + os.Getenv("METAORM_ENCRYPT_KEY") + "') ASC")
			} else {
				gormDB = gormDB.Order("AES_DECRYPT(" + query.SafeField(s.Field) + ", '" + os.Getenv("METAORM_ENCRYPT_KEY") + "') DESC")
			}
		} else {
			if s.Asc {
				gormDB = gormDB.Order(query.SafeField(s.Field) + " ASC")
			} else {
				gormDB = gormDB.Order(query.SafeField(s.Field) + " DESC")
			}
		}
	}
	return NewDB(gormDB)
}
