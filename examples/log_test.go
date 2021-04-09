package examples

import (
	"github.com/lvxin0315/gg-framework/config"
	"github.com/lvxin0315/gg-framework/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	logConfig := config.LogConfig{
		Level:    logrus.DebugLevel,
		Filename: "Debug.log",
		MaxSize:  300,
		MaxAge:   30,
	}
	l := log.NewLog(logConfig)
	_test(l)
	l.Info("hi,", " wl Info")

}

func _test(l *log.L) {
	l.Error("hi,", " wl Error")
}
