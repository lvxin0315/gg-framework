package config

import "github.com/sirupsen/logrus"

type LogConfig struct {
	// 日志级别
	Level logrus.Level
	// 文件保存位置
	Filename string
	// 文件大小
	MaxSize int
	// 最长保留时间
	MaxAge int
}
