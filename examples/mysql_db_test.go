package examples

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gg-framework/config"
	"github.com/lvxin0315/gg-framework/core"
	"testing"
	"time"
)

func TestNewMysqlDB(t *testing.T) {
	config := config.MysqlConfig{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "root",
		Port:     3306,
		Database: "demo9",
		Debug:    true,
		Charset:  "utf8mb4",
	}
	mysqlDB := core.NewMysqlDB(&config)

	for i := 0; i < 100; i++ {
		mysqlDB.Exec(func(gormDB *gorm.DB) error {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(gormDB.Table("goods").Where("id = 100").Row())
			return nil
		})
	}
}
