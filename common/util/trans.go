package util

import (
	"reflect"
	"strings"
)

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
func ObjToObj(src any, dest any, fields ...string) {
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
		if _, ok := srcType.FieldByName(destFieldType.Name); ok {
			srcFieldValue := srcValue.FieldByName(destFieldType.Name)

			// 检查字段类型是否匹配
			if srcFieldValue.Type().AssignableTo(destField.Type()) {
				// 进行赋值操作
				destField.Set(srcFieldValue)
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
