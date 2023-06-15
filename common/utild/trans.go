package utild

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ArrT[T any] []T
type MapT[k comparable, v any] map[k]v

// MapToArr 将map转成数组
func MapToArr[k comparable, v any](mp MapT[k, v]) []v {
	arr := make(ArrT[v], 0)
	var i int
	for _, data := range mp {
		arr[i] = data
		i++
	}
	return arr
}

// GetKey 泛型接口
type GetKey[k comparable] interface {
	any
	Get() k
}

var s GetKey[int]

func ArrToMap[data GetKey[k], k comparable](arr ArrT[data]) MapT[k, data] {
	mp := make(MapT[k, data], len(arr))
	for _, data := range arr {
		mp[data.Get()] = data
	}
	return mp
}

// ObjToMap 根据逗号分割的字符串，从来源结构体中获取值变成新的map
func ObjToMap(data any, str string) map[string]any {
	result := make(map[string]interface{})
	// 将逗号分隔的字符串按逗号分割为切片
	values := strings.Split(str, ",")
	// 使用反射获取传入数据的值和类型
	v := reflect.ValueOf(data)
	t := v.Type()
	for _, key := range values {
		// 检查字段是否存在于结构体类型中
		field, ok := t.FieldByName(key)
		if ok {
			// 如果字段存在，则获取字段的值并存储到结果map中
			fieldValue := v.FieldByName(key)
			result[key] = fieldValue.Interface()
		} else {
			// 如果字段不存在，则将其设为零值
			result[key] = reflect.Zero(field.Type).Interface()
		}
	}

	return result
}

// ObjToObj 根据目标结构体的字段，从来源结构体中查找对应的值并进行赋值
func ObjToObj(src interface{}, dest interface{}, fields ...string) {
	srcValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dest).Elem()

	srcType := srcValue.Type()
	destType := destValue.Type()
	// 如果存在指定的字段参数，则将其按逗号分割为字段列表
	var fieldList []string
	if len(fields) > 0 {
		fieldList = strings.Split(fields[0], ",")
	}

	// 遍历目标结构体的字段
	for i := 0; i < destValue.NumField(); i++ {
		destField := destValue.Field(i)
		destFieldType := destType.Field(i)

		// 如果存在字段参数且不在字段列表中，则跳过该字段
		if len(fieldList) > 0 && !StrInArr(destFieldType.Name, fieldList) {
			continue
		}

		// 在来源结构体中查找同名字段
		if srcFieldType, ok := srcType.FieldByName(destFieldType.Name); ok {
			srcFieldValue := srcValue.FieldByName(destFieldType.Name)

			// 检查字段类型是否匹配
			if srcFieldValue.Type().AssignableTo(destField.Type()) {
				// 进行赋值操作
				destField.Set(srcFieldValue)
			} else {
				// 如果字段类型不匹配，则输出警告信息
				fmt.Printf("Warning: field type mismatch, source field '%s' type '%s', destination field '%s' type '%s'\n",
					srcFieldType.Name, srcFieldType.Type, destFieldType.Name, destFieldType.Type)
			}
		}
	}
}

// StrInArr 判断字段是否在舒服字符串中
func StrInArr(fieldName string, fields []string) bool {
	for _, field := range fields {
		if field == fieldName {
			return true
		}
	}
	return false
}

// IpTo12Digits 将ip补0转成12个数字
func IpTo12Digits(ip string) (string, error) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return "", fmt.Errorf("invalid IP address: %s", ip)
	}
	digits := make([]string, 4)
	for i, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return "", fmt.Errorf("invalid IP address: %s", ip)
		}
		digits[i] = fmt.Sprintf("%03d", val)
	}
	return strings.Join(digits, ""), nil
}
