package log

import (
	"os"
	"wallet/pkg/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	logger *zap.Logger
}

func fromLvl(lvl Level) zapcore.Level {
	switch lvl {
	case LvlCrit:
		return zapcore.DebugLevel
	case LvlError:
		return zapcore.ErrorLevel
	case LvlWarn:
		return zapcore.WarnLevel
	case LvlInfo:
		return zapcore.InfoLevel
	case LvlDebug:
		return zapcore.DebugLevel
	case LvlTrace:
		return zapcore.DebugLevel
	}
	return zapcore.DebugLevel
}

func initInZap() {
	initInZapAt(os.Stdout, LvlInfo)
}

func InitInZapAtFile(path, prefix string, lvl Level) error {
	write, err := writeAtFile(path, prefix)
	if err != nil {
		return err
	}
	initInZapAt(zapcore.NewMultiWriteSyncer(os.Stdout, write), lvl)
	return nil
}

func initInZapAt(syncer zapcore.WriteSyncer, lvl Level) {
	encodeConf := zap.NewProductionEncoderConfig()
	encodeConf.EncodeTime = zapcore.TimeEncoderOfLayout(util.DateLayout)
	root = &zapLog{logger: zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConf),
		zapcore.Lock(syncer),
		fromLvl(lvl)),
	)}
}

func (l *zapLog) New(name string) Logger {
	return &zapLog{logger: l.logger.Named(name)}
}

func (l *zapLog) Trace(msg string, ctx ...interface{}) {
	l.logger.Sugar().Debugw(msg, ctx...)
}

func (l *zapLog) Debug(msg string, ctx ...interface{}) {
	l.logger.Sugar().Debugw(msg, ctx...)
}

func (l *zapLog) Info(msg string, ctx ...interface{}) {
	l.logger.Sugar().Infow(msg, ctx...)
}

func (l *zapLog) Warn(msg string, ctx ...interface{}) {
	l.logger.Sugar().Warnw(msg, ctx...)
}

func (l *zapLog) Error(msg string, ctx ...interface{}) {
	l.logger.Sugar().Errorw(msg, ctx...)
}
