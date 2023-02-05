package demo

import (
	"context"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

const ProblemModule = "problem"

func InitLog() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash("/home/leetcode/leetcode.log"),
		MaxSize:    50, // MB
		MaxBackups: 10,
		MaxAge:     30, // days
		Compress:   true,
	}
	logrus.SetReportCaller(true)
	logrus.SetOutput(io.MultiWriter(os.Stdout, lumberjackLogger))
}
func LogWithCtx(ctx context.Context, module string, fields ...map[string]interface{}) *logrus.Entry {
	if len(fields) > 0 {
		return logrus.WithContext(ctx).WithField("module", module).WithFields(fields[0])
	}
	return logrus.WithContext(ctx).WithField("module", module)
}

func LogModule(module string, fields ...map[string]interface{}) *logrus.Entry {
	if len(fields) > 0 {
		return logrus.WithField("module", module).WithFields(fields[0])
	}
	return logrus.WithField("module", module)
}
