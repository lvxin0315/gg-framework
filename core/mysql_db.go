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
type MysqlDB struct {
	config *config.MysqlConfig
	gormDB *gorm.DB
}

func NewMysqlDB(config *config.MysqlConfig) *MysqlDB {
	mysqlDB := new(MysqlDB)
	mysqlDB.init(config)
	return mysqlDB
}

func (mysqlDB *MysqlDB) init(config *config.MysqlConfig) {
	mysqlDB.config = config
	gormDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset))
	if err != nil {
		panic(err)
	}
	fmt.Println(gormDB)
	mysqlDB.gormDB = gormDB
	mysqlDB.gormDB.LogMode(config.Debug)

}

func (mysqlDB *MysqlDB) Exec(fs ...FuncWithDB) error {
	gormDB := mysqlDB.gormDB.New()
	defer gormDB.Close()
	for _, f := range fs {
		err := f(gormDB)
		if err != nil {
			return err
		}
	}
	return nil
}
