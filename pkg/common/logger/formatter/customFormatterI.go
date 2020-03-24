package formatter

import (
	"github.com/sirupsen/logrus"

	loggerAVA "github.com/ver13/ava/pkg/common/logger"
)

type CustomFormatterI interface {
	Format(entry *logrus.Entry) ([]byte, error)
	Sprintf(values ...interface{}) string
	parseTemplate(template string, custom CustomHandlers)
	EnableForceColors()
	IsForceColors() bool
	IsDisableColors() bool
	Color(t loggerAVA.LogLevelType) int
	SetColor(t loggerAVA.LogLevelType, reset int)
}
