package utild

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math"
)

func GetFileHashHex(file io.Reader) (string, error) {
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashHex := fmt.Sprintf("%x", hash.Sum(nil))
	return hashHex, nil
}

func FormatFileSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	// 计算文件大小的单位指数
	exp := int(math.Log(float64(size)) / math.Log(1024))
	// 将文件大小转换为指定单位的大小
	size = size / int64(math.Pow(1024, float64(exp)))

	// 根据文件大小的单位指数决定使用哪个单位
	if exp > len(units)-1 {
		exp = len(units) - 1
	}
	unit := units[exp]

	// 根据文件大小的单位指数进行舍入和格式化
	switch exp {
	case 0:
		return fmt.Sprintf("%d%s", size, unit)
	case 1, 2:
		return fmt.Sprintf("%.1f%s", float64(size), unit)
	default:
		return fmt.Sprintf("%.2f%s", float64(size), unit)
	}
}
