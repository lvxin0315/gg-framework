package examples

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gg-framework/config"
	"github.com/lvxin0315/gg-framework/core"
	"testing"
)

type TestSqliteModel struct {
	gorm.Model
	Name string
	Age  int
}

func (m *TestSqliteModel) TableName() string {
	return "test_model"
}

func TestNewSqliteDB(t *testing.T) {
	config := config.SqliteConfig{
		DBFile: "test.db",
		Debug:  true,
	}
	sqliteDB := core.NewSqliteDB(&config)

	_ = sqliteDB.Exec(func(gormDB *gorm.DB) error {
		gormDB.AutoMigrate(TestSqliteModel{})
		return nil
	})

	for i := 0; i < 100; i++ {
		sqliteDB.Exec(func(gormDB *gorm.DB) error {
			return gormDB.Model(TestSqliteModel{}).Save(&TestSqliteModel{
				Name: fmt.Sprintf("name%d", i),
				Age:  i,
			}).Error
		})
	}

	// 输出总数
	var count = 0
	sqliteDB.Exec(func(gormDB *gorm.DB) error {
		gormDB.Model(TestSqliteModel{}).Count(&count)
		return nil
	})
	fmt.Println("count:", count)
}
