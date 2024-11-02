package utild

import (
	"math/big"
)

// 64进制的字符集
const base64Charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

func UidToCode(uid string) (code string) {
	code = "0"
	// 将数字字符串转换为大整数
	decimal := new(big.Int)
	_, success := decimal.SetString(uid, 10)
	if !success {
		return
	}

	// 转换为64进制
	if decimal.Cmp(big.NewInt(0)) == 0 {
		return
	}

	base64Str := ""
	base := big.NewInt(64)

	for decimal.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		decimal.DivMod(decimal, base, mod)
		base64Str = string(base64Charset[mod.Int64()]) + base64Str
	}

	return base64Str

}
func CodeToUid(code string) (uid string) {
	uid = "0"
	base := big.NewInt(64)
	decimal := big.NewInt(0)

	for _, char := range code {
		// 找到字符在base64Charset中的索引位置
		index := int64(-1)
		for i, c := range base64Charset {
			if c == char {
				index = int64(i)
				break
			}
		}

		if index == -1 {
			return
		}

		decimal.Mul(decimal, base)
		decimal.Add(decimal, big.NewInt(index))
	}

	return decimal.String()
}
