package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvxin0315/gg-framework/config"
	"github.com/lvxin0315/gg-framework/consts"
	"github.com/lvxin0315/gg-framework/log"
)

type FuncWithMysqlGormDB func(gormDB *gorm.DB) error

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

func NewMysqlDB(config *config.MysqlConfig, logger *log.L) *MysqlDB {
	mysqlDB := new(MysqlDB)
	mysqlDB.init(config)
	//设置log
	if logger != nil {
		mysqlDB.gormDB.SetLogger(logger)
	}
	return mysqlDB
}

func (mysqlDB *MysqlDB) init(config *config.MysqlConfig) {
	mysqlDB.config = config
	mysqlDB.defaultConfig()
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
	//gormDB.DB().SetMaxIdleConns(5)
	//gormDB.DB().SetMaxOpenConns(10)
	mysqlDB.gormDB = gormDB
	mysqlDB.gormDB.LogMode(config.Debug)
}

/**
 * @Author lvxin0315@163.com
 * @Description 执行操作
 * @Date 4:49 下午 2021/4/13
 * @Param
 * @return
 **/
func (mysqlDB *MysqlDB) Exec(fs ...FuncWithMysqlGormDB) error {
	gormDB := mysqlDB.gormDB.New()
	for _, f := range fs {
		err := f(gormDB)
		if err != nil {
			return err
		}
	}
	return nil
}

/**
 * @Author lvxin0315@163.com
 * @Description 填充默认值
 * @Date 11:07 上午 2021/4/22
 * @Param
 * @return
 **/
func (mysqlDB *MysqlDB) defaultConfig() {
	if mysqlDB.config.Host == "" {
		mysqlDB.config.Host = consts.MysqlDefaultHost
	}
	if mysqlDB.config.User == "" {
		mysqlDB.config.User = consts.MysqlDefaultUser
	}
	if mysqlDB.config.Port == 0 {
		mysqlDB.config.Port = consts.MysqlDefaultPort
	}
	if mysqlDB.config.Charset == "" {
		mysqlDB.config.Charset = consts.MysqlDefaultCharset
	}
}
