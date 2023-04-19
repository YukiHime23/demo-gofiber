package repositories

import (
	"context"
	"gorm.io/gorm"
)

func SampleFunction(ctx context.Context, db *gorm.DB) []map[string]interface{} {
	var result []map[string]interface{}
	db.Table("sample").Scan(&result).WithContext(ctx)
	return result
}
