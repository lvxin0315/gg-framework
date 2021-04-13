package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lvxin0315/gg-framework/config"
)

type FuncWithSqliteGormDB func(gormDB *gorm.DB) error

/**
 * @Author lvxin0315@163.com
 * @Description sqlite操作专用
 * @Date 4:44 下午 2021/4/13
 * @Param
 * @return
 **/
type SqliteDB struct {
	config *config.SqliteConfig
	gormDB *gorm.DB
}

func NewSqliteDB(config *config.SqliteConfig) *SqliteDB {
	sqliteDB := new(SqliteDB)
	sqliteDB.init(config)
	return sqliteDB
}

func (sqliteDB *SqliteDB) init(config *config.SqliteConfig) {
	sqliteDB.config = config
	gormDB, err := gorm.Open("sqlite3", config.DBFile)
	if err != nil {
		panic(err)
	}
	sqliteDB.gormDB = gormDB
	sqliteDB.gormDB.LogMode(config.Debug)
}

/**
 * @Author lvxin0315@163.com
 * @Description 执行操作
 * @Date 4:49 下午 2021/4/13
 * @Param
 * @return
 **/
func (sqliteDB *SqliteDB) Exec(fs ...FuncWithSqliteGormDB) error {
	gormDB := sqliteDB.gormDB.New()
	for _, f := range fs {
		err := f(gormDB)
		if err != nil {
			return err
		}
	}
	return nil
}
