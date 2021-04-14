package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

/**
 * @Author lvxin0315@163.com
 * @Description gorm带uuid
 * @Date 3:46 下午 2021/4/14
 * @Param
 * @return
 **/
type GormWithUUID struct {
	UUID string `gorm:"column:uuid;type:varchar(255);not null"`
}

func (m *GormWithUUID) BeforeCreate(tx *gorm.DB) (err error) {
	m.UUID = uuid.NewV4().String()
	return
}
