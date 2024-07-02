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

// Info 打印日志到控制台
func Info(content string) {
	// 获取调用者的信息
	_, file, line, ok := runtime.Caller(1)
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
