package dao

import (
	"database/sql"
	"fmt"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"reflect"
	"strings"
)

// PrepareData 根据结构体转成model使用的map
func PrepareData(source interface{}) (map[string]any, error) {
	sourceValues, err := processSource(source)
	if err != nil {
		return nil, err
	}
	//获取目标结构体的属性集合
	sourceTypes := sourceValues.Type()
	//遍历目标结构体所有字段
	result := make(map[string]any)
	for i := 0; i < sourceValues.NumField(); i++ {
		//获取结构体的值对象
		field := sourceValues.Field(i)
		sourceName := sourceTypes.Field(i).Name
		//将需要提取的字段也转成蛇形
		targetName := utild.StrToSnake(sourceName)
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
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(sql.NullString{}) {
				ns := field.Interface().(sql.NullString)
				if ns.Valid {
					result[targetName] = ns.String
				} else {
					result[targetName] = ""
				}
			}

		}
	}
	return result, nil
}

// PrepareDataByTarget 根据结构体、目标字段生成写入model的数据
func PrepareDataByTarget(source interface{}, targetDelimiterSeparated string, isEmptySet ...bool) (map[string]any, error) {
	if targetDelimiterSeparated == "" {
		return nil, resd.NewErr("PrepareDataByTarget中targetDelimiterSeparated参数为空")
	}
	targetFields := strings.Split(targetDelimiterSeparated, ",")
	sourceValues, err := processSource(source)
	if err != nil {
		return nil, err
	}
	//获取目标字段的集合
	targets := make(map[string]int)
	for _, v := range targetFields {
		//将结构体里的字段转成蛇形，最终都按蛇形匹配
		targets[utild.StrToSnake(v)] = 0
	}
	//获取目标结构体的属性集合
	sourceTypes := sourceValues.Type()
	//遍历目标结构体所有字段
	result := make(map[string]any)
	for i := 0; i < sourceValues.NumField(); i++ {
		//获取结构体的值对象
		field := sourceValues.Field(i)
		sourceName := sourceTypes.Field(i).Name
		//将需要提取的字段也转成蛇形
		targetName := utild.StrToSnake(sourceName)
		//判断是否在目标字段中，不是就跳过
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

// PrepareDataByExcept 根据结构体、目标字段生成写入model的数据
func PrepareDataByExcept(source interface{}, exceptDelimiterSeparated string) (map[string]string, error) {
	if exceptDelimiterSeparated == "" {
		return nil, resd.NewErr("PrepareDataByExcept中exceptDelimiterSeparated参数为空")
	}
	exceptFields := strings.Split(exceptDelimiterSeparated, ",")
	sourceValues, err := processSource(source)
	if err != nil {
		return nil, err
	}
	//获取目标字段的集合
	excepts := make(map[string]int)
	for _, v := range exceptFields {
		//将结构体里的字段转成蛇形，最终都按蛇形匹配
		excepts[utild.StrToSnake(v)] = 0
	}
	//获取目标结构体的属性集合
	sourceTypes := sourceValues.Type()
	//遍历目标结构体所有字段
	result := make(map[string]string)
	for i := 0; i < sourceValues.NumField(); i++ {
		//获取结构体的值对象
		field := sourceValues.Field(i)
		sourceName := sourceTypes.Field(i).Name
		//将需要提取的字段也转成蛇形
		targetName := utild.StrToSnake(sourceName)
		//判断是否在目标字段中，不是就跳过
		if _, ok := excepts[targetName]; ok {
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
	}

	return result, nil
}
func processSource(source interface{}) (reflect.Value, error) {
	sourceValues := reflect.ValueOf(source)
	//如果是指针类型，继续获取判断
	if sourceValues.Kind() == reflect.Ptr {
		sourceElem := sourceValues.Elem()
		if sourceElem.Kind() == reflect.Struct {
			sourceValues = sourceElem
		} else {
			return sourceValues, resd.NewErr("指针类型的source必须为结构体的指针")
		}
	}
	if sourceValues.Kind() != reflect.Struct {
		return sourceValues, resd.NewErr("source必须为结构体或结构体指针")
	}
	return sourceValues, nil
}
