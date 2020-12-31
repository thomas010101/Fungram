package log

import (
	"fmt"
	"io"
	"os"
	"time"
	"wallet/pkg/util"
)

type Logger interface {
	New(name string) Logger
	Trace(msg string, ctx ...interface{})
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
}

type Level int

const (
	LvlCrit Level = iota
	LvlError
	LvlWarn
	LvlInfo
	LvlDebug
	LvlTrace
)

func (l Level) Decode(s string) (interface{}, error) {
	switch s {
	case "trace", "trce":
		return LvlTrace, nil
	case "debug", "dbug":
		return LvlDebug, nil
	case "info":
		return LvlInfo, nil
	case "warn":
		return LvlWarn, nil
	case "error", "eror":
		return LvlError, nil
	case "crit":
		return LvlCrit, nil
	default:
		return LvlDebug, fmt.Errorf("unknown level: %v", s)
	}
}

var (
	root Logger
)

func init() {
	initInZap()
}

func New(name string) Logger {
	return root.New(name)
}

func Trace(msg string, ctx ...interface{}) {
	root.Trace(msg, ctx...)
}

func Debug(msg string, ctx ...interface{}) {
	root.Debug(msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	root.Info(msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	root.Warn(msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	root.Error(msg, ctx...)
}

type WriteSyncer interface {
	io.Writer
	Sync() error
}

func writeAtFile(path, prefix string) (WriteSyncer, error) {
	var name string
	if len(prefix) > 0 {
		name = fmt.Sprintf("%s-%s", prefix, util.FormatDate(time.Now(), util.NotSeparatorDataLayout))
	} else {
		name = util.FormatDate(time.Now(), util.NotSeparatorDataLayout)
	}
	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s.log", path, name),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0600)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	return file, nil
}
