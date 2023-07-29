package utild

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func GetFileSha1ByOsFile(file *os.File) (string, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	sha1Hash := sha1.New()
	if _, err := io.Copy(sha1Hash, file); err != nil {
		return "", err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(sha1Hash.Sum(nil)), nil
}
func GetFileSha1ByIoReader(file io.Reader) (string, error) {
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashHex := fmt.Sprintf("%x", hash.Sum(nil))
	return hashHex, nil
}

func FormatFileSize(size int64) string {
	// 文件大小的单位列表
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	// 计算文件大小的单位指数
	exp := int(math.Log(float64(size)) / math.Log(1024)) // 计算文件大小的单位指数
	expSize := int64(math.Pow(1024, float64(exp)))       // 计算当前单位的大小
	if expSize == 0 {
		return fmt.Sprintf("%d", expSize) // 文件大小为 0
	}
	// 将文件大小转换为指定单位的大小
	floatSize := float64(size) / float64(expSize)

	// 根据文件大小的单位指数决定使用哪个单位
	if exp > len(units)-1 {
		exp = len(units) - 1 // 文件大小超出最大单位，使用最大单位
	}
	unit := units[exp] // 获取当前单位

	// 根据文件大小的单位指数进行舍入和格式化
	s := ""
	switch exp {
	case 0:
		s = fmt.Sprintf("%f", floatSize) // 文件大小小于 1KB，直接输出整数部分和单位
	case 1:
		s = fmt.Sprintf("%.1f", floatSize) // 文件大小在 1KB 和 1MB 之间，保留一位小数
	case 2, 3, 4:
		s = fmt.Sprintf("%.2f", floatSize) // 文件大小在 1MB 和 1TB 之间，保留两位小数
	default:
		s = fmt.Sprintf("%.3f", floatSize) // 文件大小在 1TB 和最大单位之间，保留三位小数
	}
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	return fmt.Sprintf("%s%s", s, unit)
}
