package logger

import (
	"fmt"
	"io"
	"log"
)

const (
	LvDebug = iota
	LvInfo
	LvWarn
	LvError
	LvFatal
)

var (
	gLogger  *log.Logger
	gLevel   = LvDebug
	gLogPrefix = []string{
		"[Debug] ",
		"[Info] ",
		"[Warn] ",
		"[Error] ",
		"[Fatal] ",
	}
)

func InitFileLog(out io.Writer, appName string, l int) {
	if len(appName) <= 0 {
		appName = "DefaultApp"
	}

	gLogger = log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile)
	if gLogger == nil {
		panic("InitFileLog log.New == nil")
		return
	}

	if l > LvFatal || l < LvDebug {
		gLevel = LvDebug
	} else {
		gLevel = l
	}
}

func _log(lv int, format string, v ...interface{}) {
	if lv < gLevel {
		return
	}
	str := fmt.Sprintf(gLogPrefix[lv]+format, v...)
	gLogger.Output(3, str)
}

func Debug(format string, v ...interface{}) { _log(LvDebug, format, v...) }
func Info(format string, v ...interface{})  { _log(LvInfo, format, v...) }
func Warn(format string, v ...interface{})  { _log(LvWarn, format, v...) }
func Error(format string, v ...interface{}) { _log(LvError, format, v...) }
func Fatal(format string, v ...interface{}) {
	_log(LvFatal, format, v...)
	panic(fmt.Sprintf(format, v...))
}
