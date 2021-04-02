package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvxin0315/gg-framework/config"
)

type FuncWithDB func(gormDB *gorm.DB) error

/**
 * @Author lvxin0315@163.com
 * @Description mysql操作专用
 * @Date 2:45 下午 2021/4/2
 * @Param
 * @return
 **/
func MysqlDB(config *config.MysqlConfig, fs ...FuncWithDB) error {
	gormDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset))
	if err != nil {
		return err
	}
	gormDB.LogMode(config.Debug)
	defer gormDB.Close()
	for _, f := range fs {
		err = f(gormDB)
		if err != nil {
			return err
		}
	}
	return nil
}
