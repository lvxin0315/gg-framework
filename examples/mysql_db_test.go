package examples

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gg-framework/config"
	"github.com/lvxin0315/gg-framework/core"
	"github.com/lvxin0315/gg-framework/log"
	"github.com/lvxin0315/gg-framework/model"
	"github.com/sirupsen/logrus"
	"testing"
)

type TestMysqlModel struct {
	gorm.Model
	model.GormWithUUID
	Name string
	Age  int
}

func TestNewMysqlDB(t *testing.T) {
	logConfig := config.LogConfig{
		Level:    logrus.DebugLevel,
		Filename: "Debug.log",
		MaxSize:  300,
		MaxAge:   30,
	}
	l := log.NewLog(logConfig)

	config := config.MysqlConfig{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "root",
		Port:     3306,
		Database: "demo9",
		Debug:    true,
		Charset:  "utf8mb4",
	}
	mysqlDB := core.NewMysqlDB(&config, l)

	_ = mysqlDB.Exec(func(gormDB *gorm.DB) error {
		gormDB.AutoMigrate(TestMysqlModel{})
		return nil
	})

	for i := 0; i < 100; i++ {
		mysqlDB.Exec(func(gormDB *gorm.DB) error {
			return gormDB.Model(TestMysqlModel{}).Save(&TestMysqlModel{
				Name: fmt.Sprintf("name%d", i),
				Age:  i,
			}).Error
		})
	}

	// 输出总数
	var count = 0
	mysqlDB.Exec(func(gormDB *gorm.DB) error {
		gormDB.Model(TestMysqlModel{}).Count(&count)
		return nil
	})
	fmt.Println("count:", count)

}
