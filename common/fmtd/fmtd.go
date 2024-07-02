package fmtd

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorReset   = "\033[0m"
)

// Logger 结构体包含日志相关的信息
type Logger struct {
	callerDepth int
}

// WithCaller 设置调用者信息的深度
func WithCaller(depth int) *Logger {
	return &Logger{callerDepth: depth}
}

// Info 打印日志到控制台
func (l *Logger) Info(content string) {
	// 获取调用者的信息
	_, file, line, ok := runtime.Caller(l.callerDepth + 1)
	if !ok {
		log.Println("Failed to get caller info")
		return
	}

	// 获取当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 打印带有颜色的日志
	fmt.Printf("%s[Info]%s%s%s%s %s%s:%d%s %s%s%s\n",
		ColorRed, ColorReset,
		ColorCyan, currentTime, ColorReset,
		ColorMagenta, file, line, ColorReset,
		ColorGreen, content, ColorReset)
}

// Info 方法，使用默认的调用者深度
func Info(content string) {
	WithCaller(1).Info(content)
}
