package log

import (
	"github.com/lvxin0315/gg-framework/config"
	"github.com/sirupsen/logrus"
)

var l *L

var defaultLogConfig = config.LogConfig{
	Level:    logrus.InfoLevel,
	Filename: "gg.log",
	MaxSize:  DefaultLogMaxSize,
	MaxAge:   DefaultLogMaxAge,
}

func InitLog() {
	l = NewLog(defaultLogConfig)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Warn(args ...interface{}) {
	l.Warn(args...)
}

func Fatal(args ...interface{}) {
	l.Fatal(args...)
}
