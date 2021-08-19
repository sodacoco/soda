package util

import (
	"fmt"
	"os"
	"time"
)

const (
	// 错误
	LevelError = iota
	// 警告
	LevelWarning
	// 提示
	LevelInformational
	// 排错
	LevelDebug
)

var logger *Logger

type Logger struct {
	level int
}

// 打印
func (ll *Logger) Println(msg string) {
	_, _ = fmt.Printf("%s %s\n", time.Now().Format("2006-01-02 15:04:05 -0700"), msg)
}

// 极端错误
func (ll *Logger) Panic(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[Panic] "+format, v...)
	ll.Println(msg)
	os.Exit(0)
}

// 错误
func (ll *Logger) Error(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[E] "+format, v...)
	ll.Println(msg)
}

// 警告
func (ll *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > ll.level {
		return
	}
	msg := fmt.Sprintf("[W] "+format, v...)
	ll.Println(msg)
}

// 信息
func (ll *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > ll.level {
		return
	}
	msg := fmt.Sprintf("[I] "+format, v...)
	ll.Println(msg)
}

// 构建logger
func BuildLogger(level string) {
	intLevel := LevelError
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInformational
	case "debug":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	logger = &l
}

// 返回日志对象
func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}
	return logger
}
