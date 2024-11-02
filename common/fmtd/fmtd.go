package fmtd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

// logger 结构体包含日志相关的信息
type logger struct {
	callerDepth int
}

// Info 打印日志到控制台
func (l *logger) Info(content ...any) {
	for _, v := range content {
		l.print(fmt.Sprintf("%v", v), "info")
	}
}

// Json 转成json打印
func (l *logger) Json(data any) {
	v, _ := json.Marshal(data)
	l.print(string(v), "info")
}

// Error 打印日志到控制台
func (l *logger) Error(content ...any) {
	for _, v := range content {
		l.print(fmt.Sprintf("%v", v), "error")
	}
}
func (l *logger) print(content string, level string) {
	// 获取调用者的信息
	_, file, line, ok := runtime.Caller(l.callerDepth + 1)
	if !ok {
		log.Println("Failed to get caller info")
		return
	}

	// 获取当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// 打印带有颜色的日志
	if level == "error" {
		fmt.Printf("%s[error]%s%s%s%s %s%s:%d%s %s%s%s\n",
			ColorRed, ColorReset,
			ColorCyan, currentTime, ColorReset,
			ColorMagenta, file, line, ColorReset,
			ColorGreen, content, ColorReset)
	} else {
		fmt.Printf("%s[info]%s%s%s%s %s%s:%d%s %s%s%s\n",
			ColorYellow, ColorReset,
			ColorCyan, currentTime, ColorReset,
			ColorMagenta, file, line, ColorReset,
			ColorGreen, content, ColorReset)
	}

}

// Info 方法，使用默认的调用者深度
func Info(content ...any) {
	WithCaller(2).Info(content...)
}

// Error 方法，使用默认的调用者深度
func Error(content ...any) {
	WithCaller(2).Error(content...)
}
func Json(data any) {
	WithCaller(2).Json(data)
}
func Fatal(content ...any) {
	WithCaller(2).Error(content...)
	os.Exit(1)
}

// WithCaller 设置调用者信息的深度
func WithCaller(depth int) *logger {
	return &logger{callerDepth: depth}
}
