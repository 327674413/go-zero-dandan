package utild

import (
	"errors"
	"fmt"
	"go-zero-dandan/common/resd"
	"reflect"
	"strings"
)

func MakeModelData(source interface{}, targetDelimiterSeparated string, isEmptySet ...bool) (map[string]string, error) {
	if targetDelimiterSeparated == "" {
		return nil, resd.NewErr("MakeModelData中targetDelimiterSeparated参数为空")
	}
	fields := strings.Split(targetDelimiterSeparated, ",")
	sourceValues := reflect.ValueOf(source)
	//如果是指针类型，继续获取判断
	if sourceValues.Kind() == reflect.Ptr {
		sourceElem := sourceValues.Elem()
		if sourceElem.Kind() == reflect.Struct {
			sourceValues = sourceElem
		} else {
			return nil, errors.New("source must be a struct point")
		}
	}
	if sourceValues.Kind() != reflect.Struct {
		return nil, errors.New("source must be a struct")
	}
	//获取目标字段的集合
	targets := make(map[string]int)
	for _, v := range fields {
		//将结构体里的字段转成蛇形，最终都按蛇形匹配
		targets[StrToSnake(v)] = 0
	}
	result := make(map[string]string)
	//获取目标结构体的属性集合
	sourceTypes := sourceValues.Type()
	fmt.Println(targets)
	//遍历目标结构体所有字段
	for i := 0; i < sourceValues.NumField(); i++ {
		//获取结构体的值对象
		field := sourceValues.Field(i)
		sourceName := sourceTypes.Field(i).Name
		//将需要提取的字段也转成蛇形
		targetName := StrToSnake(sourceName)
		if _, ok := targets[targetName]; !ok {
			continue
		}
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result[targetName] = fmt.Sprintf("%d", field.Int())
		case reflect.String:
			result[targetName] = field.String()
		case reflect.Ptr:
			if !field.IsNil() {
				prtVal := field.Elem()
				switch prtVal.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					result[targetName] = fmt.Sprintf("%d", prtVal.Int())
				case reflect.String:
					result[targetName] = prtVal.String()
				}
			}

		}
		targets[sourceName] = 1
	}
	if len(isEmptySet) > 0 && isEmptySet[0] {
		for k, v := range targets {
			if v == 0 {
				result[k] = ""
			}
		}
	}

	return result, nil
}
func SetModelFromMap(source map[string]string, targetModel any) error {
	targetValuePt := reflect.ValueOf(targetModel)
	if targetValuePt.Kind() != reflect.Ptr {
		return errors.New("target must a struct point")
	}
	targetStv := targetValuePt.Elem()
	if targetStv.Kind() != reflect.Struct {
		return errors.New("target must a struct point")
	}
	targetTypePt := reflect.TypeOf(targetModel)
	targetStk := targetTypePt.Elem()
	for i := 0; i < targetStv.NumField(); i++ {
		field := targetStk.Field(i)
		if !targetStv.Field(i).CanSet() {
			continue
		}
		if val, ok := source[StrToSnake(field.Name)]; ok {
			switch field.Type.Kind() {
			case reflect.String:
				targetStv.Field(i).SetString(val)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				targetStv.Field(i).SetInt(AnyToInt64(val))
			}
		}
	}
	return nil
}
