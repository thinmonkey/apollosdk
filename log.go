package apollosdk

import (
	"github.com/cihub/seelog"
)

var Logger LoggerInterface

func init() {
	initLogger(initSeeLog("seelog.xml"))
}

func initLogger(ILogger LoggerInterface) {
	Logger = ILogger
}

type LoggerInterface interface {
	Debugf(format string, params ...interface{})

	Infof(format string, params ...interface{})

	Warnf(format string, params ...interface{}) error

	Errorf(format string, params ...interface{}) error

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{}) error

	Error(v ...interface{}) error
}

func initSeeLog(configPath string) (LoggerInterface) {
	logger, err := seelog.LoggerFromConfigAsFile(configPath)

	//if error is happen change to default config.
	if err != nil {
		logger, err = seelog.LoggerFromConfigAsBytes([]byte("<seelog />"))
	}

	logger.SetAdditionalStackDepth(1)
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	return logger
}
