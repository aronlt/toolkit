package toolkit

import (
	"context"

	"github.com/sirupsen/logrus"
)

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
