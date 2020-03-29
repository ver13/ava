package formatter

import (
	"github.com/sirupsen/logrus"

	"github.com/ver13/ava/pkg/common/logger"
)

type CustomFormatterI interface {
	Format(entry *logrus.Entry) ([]byte, error)
	Sprintf(values ...interface{}) string
	parseTemplate(template string, custom CustomHandlers)
	EnableForceColors()
	IsForceColors() bool
	IsDisableColors() bool
	Color(t logger.LogLevelType) int
	SetColor(t logger.LogLevelType, reset int)
}
