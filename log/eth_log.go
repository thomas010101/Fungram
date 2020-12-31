package log

import (
	"fmt"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/log"
)

type ethLogger struct {
	moduleName string
	logger     log.Logger
}

func initInEth() {
	eLogger := log.Root()
	eLogger.SetHandler(log.StdoutHandler)
	root = &ethLogger{
		moduleName: "",
		logger:     eLogger,
	}
}

func InitInEthAtFile(path, prefix string, lvl Level) error {
	write, err := writeAtFile(path, prefix)
	if err != nil {
		return err
	}
	gLogger := log.NewGlogHandler(log.StreamHandler(io.MultiWriter(os.Stdout, write), log.LogfmtFormat()))
	gLogger.Verbosity(log.Lvl(lvl))
	eLogger := log.Root()
	eLogger.SetHandler(gLogger)
	root = &ethLogger{
		moduleName: "",
		logger:     eLogger,
	}
	return nil
}

func (l *ethLogger) New(name string) Logger {
	if len(l.moduleName) > 0 {
		name = fmt.Sprintf("%s:%s", l.moduleName, name)
	}
	tl := l.logger.New("module", name)
	return &ethLogger{logger: tl, moduleName: name}
}

func (l *ethLogger) Trace(msg string, ctx ...interface{}) {
	l.logger.Trace(msg, ctx...)
}

func (l *ethLogger) Debug(msg string, ctx ...interface{}) {
	l.logger.Debug(msg, ctx...)
}

func (l *ethLogger) Info(msg string, ctx ...interface{}) {
	l.logger.Info(msg, ctx...)
}

func (l *ethLogger) Warn(msg string, ctx ...interface{}) {
	l.logger.Warn(msg, ctx...)
}

func (l *ethLogger) Error(msg string, ctx ...interface{}) {
	l.logger.Error(msg, ctx...)
}
