package tool

import "gorm.io/gorm"

type AppContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *AppContext {
	return &AppContext{
		db: db,
	}
}

func (aCtx *AppContext) GetDBConnection() *gorm.DB {
	return aCtx.db
}
