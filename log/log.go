package log

import (
	"strings"

	"github.com/vishal1132/bootstrap-go-web/config"
	"github.com/vishal1132/bootstrap-go-web/utils"
	"go.uber.org/zap"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

func SetupLogger(config *config.LogConfig) {
	defaultLevel := zap.InfoLevel
	switch strings.ToUpper(config.Level) {
	case DEBUG:
		defaultLevel = zap.DebugLevel
	case INFO:
		defaultLevel = zap.InfoLevel
	case WARN:
		defaultLevel = zap.WarnLevel
	case ERROR:
		defaultLevel = zap.ErrorLevel
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(defaultLevel)
	zapConfig.Sampling = nil
	zapConfig.DisableStacktrace = true
	l := utils.Must(zapConfig.Build(zap.AddCallerSkip(1)))
	zap.ReplaceGlobals(l)
}
