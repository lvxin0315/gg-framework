package log

import (
	"fmt"
	"github.com/lvxin0315/gg-framework/config"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
)

const (
	DefaultLogMaxSize = 5
	DefaultLogMaxAge  = 30
)

func NewLog(config config.LogConfig) *L {
	l := new(L)
	l.initFileLogger(config)
	l.initConsoleEntry(config)
	return l
}

func NewLogDefault(level logrus.Level, logName string) *L {
	defaultConfig := config.LogConfig{
		Level:    level,
		Filename: logName,
		MaxSize:  DefaultLogMaxSize,
		MaxAge:   DefaultLogMaxAge,
	}
	return NewLog(defaultConfig)
}

/**
 * @Author lvxin0315@163.com
 * @Description 日志处理
 * @Date 1:46 下午 2021/4/7
 * @Param
 * @return
 **/
type L struct {
	// 文件输出
	fileLogger *logrus.Logger
	fileEntry  *logrus.Entry
	// 控制台输出
	consoleLogger *logrus.Logger
	consoleEntry  *logrus.Entry
}

func (l *L) initFileLogger(config config.LogConfig) {
	fileLogger := logrus.New()
	fileLogger.SetLevel(config.Level)
	fileLogger.SetOutput(&lumberjack.Logger{
		Filename: config.Filename,
		MaxSize:  config.MaxSize,
		MaxAge:   config.MaxAge,
	})
	l.fileLogger = fileLogger
	l.fileEntry = logrus.NewEntry(l.fileLogger)
}

func (l *L) initConsoleEntry(config config.LogConfig) {
	consoleLogger := logrus.New()
	consoleLogger.SetLevel(config.Level)
	consoleLogger.SetOutput(os.Stdout)
	l.consoleLogger = consoleLogger
	l.consoleEntry = logrus.NewEntry(l.consoleLogger)
}

func (l *L) Trace(args ...interface{}) {
	l.consoleEntry.Trace(args...)
	l.fileEntry.Trace(args...)
}

func (l *L) Debug(args ...interface{}) {
	l.consoleEntry.Debug(args...)
	l.fileEntry.Debug(args...)
}

func (l *L) Print(args ...interface{}) {
	l.consoleEntry.Print(args...)
	l.fileEntry.Print(args...)
}

func (l *L) Info(args ...interface{}) {
	l.consoleEntry.Info(args...)
	l.fileEntry.Info(args...)
}

func (l *L) Warn(args ...interface{}) {
	l.consoleEntry.Warn(args...)
	l.fileEntry.Warn(args...)
}

func (l *L) Warning(args ...interface{}) {
	l.consoleEntry.Warning(args...)
	l.fileEntry.Warning(args...)
}

func (l *L) Error(args ...interface{}) {
	l.consoleEntry.Error(args...)
	l.fileEntry.Error(args...)
	l._debugInfo()
}

func (l *L) Fatal(args ...interface{}) {
	l.consoleEntry.Fatal(args...)
	l.fileEntry.Fatal(args...)
}

func (l *L) Tracef(format string, args ...interface{}) {
	l.consoleEntry.Trace(fmt.Sprintf(format, args...))
	l.fileEntry.Trace(fmt.Sprintf(format, args...))
}

func (l *L) Debugf(format string, args ...interface{}) {
	l.consoleEntry.Debug(fmt.Sprintf(format, args...))
	l.fileEntry.Debug(fmt.Sprintf(format, args...))
}

func (l *L) Infof(format string, args ...interface{}) {
	l.consoleEntry.Info(fmt.Sprintf(format, args...))
	l.fileEntry.Info(fmt.Sprintf(format, args...))
}

func (l *L) Printf(format string, args ...interface{}) {
	l.consoleEntry.Print(fmt.Sprintf(format, args...))
	l.fileEntry.Print(fmt.Sprintf(format, args...))
}

func (l *L) Warnf(format string, args ...interface{}) {
	l.consoleEntry.Warn(fmt.Sprintf(format, args...))
	l.fileEntry.Warn(fmt.Sprintf(format, args...))
}

func (l *L) Warningf(format string, args ...interface{}) {
	l.consoleEntry.Warning(fmt.Sprintf(format, args...))
	l.fileEntry.Warning(fmt.Sprintf(format, args...))
}

func (l *L) Errorf(format string, args ...interface{}) {
	l.consoleEntry.Error(fmt.Sprintf(format, args...))
	l.fileEntry.Error(fmt.Sprintf(format, args...))
	l._debugInfo()
}

func (l *L) Fatalf(format string, args ...interface{}) {
	l.consoleEntry.Fatal(fmt.Sprintf(format, args...))
	l.fileEntry.Fatal(fmt.Sprintf(format, args...))
}

func (l *L) Panicf(format string, args ...interface{}) {
	l.consoleEntry.Panic(fmt.Sprintf(format, args...))
	l.fileEntry.Panic(fmt.Sprintf(format, args...))
}

/**
 * @Author lvxin0315@163.com
 * @Description 显示debug信息
 * @Date 4:46 下午 2021/4/9
 * @Param
 * @return
 **/
func (l *L) _debugInfo() {
	debug.PrintStack()
	// 分行显示
	//for _, stack := range bytes.Split(debug.Stack(), []byte("\n")) {
	//	l.fileEntry.Error(string(bytes.ReplaceAll(stack, []byte("\t"), []byte(" "))))
	//}
	l.fileLogger.Out.Write(debug.Stack())
}
