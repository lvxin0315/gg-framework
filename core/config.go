package core

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var ggViper *viper.Viper

/**
 * @Author lvxin0315@163.com
 * @Description 加载配置文件
 * @Date 2:24 下午 2021/4/2
 * @Param
 * @return
 **/
func InitConfig(configPath string, configStructList ...interface{}) error {
	ggViper = viper.New()
	ggViper.SetConfigType("toml")
	ggViper.AddConfigPath(configPath)
	err := ggViper.ReadInConfig()
	if err != nil {
		return err
	}
	// 解析
	for _, configItem := range configStructList {
		err = ggViper.Unmarshal(configItem)
		if err != nil {
			return err
		}
	}
	return nil
}
