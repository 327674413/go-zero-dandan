package utild

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 使用sha-256加密
func Sha256(str string) string {
	// 创建 SHA-256 哈希对象
	hash := sha256.New()
	// 将字符串转换为字节数组，并计算哈希值
	hash.Write([]byte(str))
	sha256Hash := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串表示
	sha256Str := hex.EncodeToString(sha256Hash)
	return sha256Str
}

// Sha1 使用sha1加密
func Sha1(str string) string {
	// 创建 SHA-256 哈希对象
	hash := sha1.New()
	// 将字符串转换为字节数组，并计算哈希值
	hash.Write([]byte(str))
	sha256Hash := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串表示
	sha256Str := hex.EncodeToString(sha256Hash)
	return sha256Str
}

// Sha1Byte 使用sha1加密Byte
func Sha1Byte(byte []byte) string {
	// 创建 SHA-256 哈希对象
	hash := sha1.New()
	// 将字符串转换为字节数组，并计算哈希值
	hash.Write(byte)
	sha256Hash := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串表示
	sha256Str := hex.EncodeToString(sha256Hash)
	return sha256Str
}
