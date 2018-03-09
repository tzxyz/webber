package webber

import (
	"runtime"
	"path/filepath"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	logger.Level = logrus.DebugLevel
	logger.Hooks.Add(&TraceHook{})
}

type TraceHook struct{}

func (h *TraceHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (h *TraceHook) Fire(entry *logrus.Entry) error {
	pc, file, line, ok := runtime.Caller(6)
	if ok {
		fn := runtime.FuncForPC(pc)
		entry.Data["func"] = fn.Name()
		entry.Data["file"] = filepath.Base(file)
		entry.Data["line"] = line
	}
	return nil
}

func Debug(args ...interface{})  {
	logger.Debug(args)
}

func Info(args ...interface{})  {
	logger.Info(args)
}

func Warn(args ...interface{})  {
	logger.Warn(args)
}

func Error(args ...interface{})  {
	logger.Error(args)
}

func Panic(args ...interface{})  {
	logger.Panic(args)
}

func Fatal(args ...interface{})  {
	logger.Fatal(args)
}