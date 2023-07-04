package main

import (
	"errors"
	"fmt"
	"go-zero-dandan/common/utild"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	/*source := map[string]string{
		"id":        "12324321312",
		"name":      "张三",
		"create_at": "19212422211",
		"update_at": "0",
	}
	targetObj := &struct {
		Id       int64
		Name     string
		CreateAt int64
		Sex      int64
	}{Id: 123}
	MapStrToStruct(source, targetObj)
	fmt.Println(targetObj)*/
	createAt := int64(5)
	source := &struct {
		Id       int64
		Name     string
		CreateAt *int64
		Sex      *int64
		EmptyStr string
	}{Id: 123, Name: "张三", CreateAt: &createAt}

	fmt.Println(StructToStrMapFrom3(source, "Id,Name,CreateAt,Sex,else,EmptyStr"))
}
func StructToStrMapFrom3(source interface{}, targetDelimiterSeparated string, isEmptySet ...bool) (map[string]string, error) {
	if targetDelimiterSeparated == "" {
		return nil, errors.New("delimiter separated target is empty")
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
		targets[v] = 0
	}
	result := make(map[string]string)
	//获取目标结构体的属性集合
	sourceTypes := sourceValues.Type()
	//遍历目标结构体所有字段
	for i := 0; i < sourceValues.NumField(); i++ {
		//获取结构体的值对象
		field := sourceValues.Field(i)
		sourceName := sourceTypes.Field(i).Name
		targetName := utild.StrToSnake(sourceName)
		if _, ok := targets[sourceName]; !ok {
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
func MapStrToStruct(source map[string]string, targetObj any) error {
	targetValuePt := reflect.ValueOf(targetObj)
	if targetValuePt.Kind() != reflect.Ptr {
		return errors.New("target must a struct point")
	}
	targetStv := targetValuePt.Elem()
	if targetStv.Kind() != reflect.Struct {
		return errors.New("target must a struct point")
	}
	targetTypePt := reflect.TypeOf(targetObj)
	targetStk := targetTypePt.Elem()
	for i := 0; i < targetStv.NumField(); i++ {
		field := targetStk.Field(i)
		if !targetStv.Field(i).CanSet() {
			continue
		}
		if val, ok := source[utild.StrToSnake(field.Name)]; ok {
			switch field.Type.Kind() {
			case reflect.String:
				targetStv.Field(i).SetString(val)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				targetStv.Field(i).SetInt(utild.AnyToInt64(val))
			}
		}
	}
	return nil
}

/*
func testSnow() {
	start := time.Now()
	target := make(map[int64]int)
	for i := 0; i < 100000; i++ {
		id := utild.MakeId()
		if _, ok := target[id]; ok {
			target[id] = target[id] + 1
		} else {
			target[id] = 0
		}
	}
	end := time.Since(start)
	fmt.Println("键数：", len(target), "，耗时：", end)
}
*/

func StructToStrMapFrom2(obj interface{}, fields ...string) map[string]string {
	result := make(map[string]string)
	r := reflect.ValueOf(obj).Elem()
	for _, field := range fields {
		f := r.FieldByName(field)
		if f.IsValid() {
			var value string
			if f.Kind() == reflect.Ptr && f.Type().Elem().Kind() != reflect.Uint8 {
				if !f.IsNil() {
					if v := f.Elem().Interface(); v != nil {
						if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", reflect.Zero(f.Type().Elem()).Interface()) {
							value = fmt.Sprintf("%v", v)
						}
					}
				}
			} else {
				if v := f.Interface(); v != nil {
					if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", reflect.Zero(f.Type()).Interface()) {
						value = fmt.Sprintf("%v", v)
					}
				}
			}
			if value != "" {
				result[field] = value
			}
		}
	}
	return result
}
func StructToStrMapFrom(obj any, targetField ...string) map[string]string {
	result := make(map[string]string)
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()
	if objType.Kind() != reflect.Struct {
		return result
	}
	for i := 0; i < objType.NumField(); i++ {
		field := objValue.Field(i)
		kind := objType.Field(i)
		fieldName := objType.Field(i).Name
		if StrInArr(fieldName, targetField) {
			v, err := getReflectValueToStr(field, kind)
			fmt.Println(fieldName, v, err)
			if err == nil {
				result[fieldName] = v
			}
		}
	}
	return result
}
func getReflectValueToStr(field reflect.Value, kind reflect.StructField) (string, error) {
	valueKind := field.Kind()
	switch valueKind {
	case reflect.String:
		return field.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(field.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(field.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'f', -1, 64), nil
	case reflect.Ptr:
		fmt.Println("验证：", field, field.IsValid(), !field.IsNil(), field.Elem())
		if field.IsValid() && !field.IsNil() {
			v := field.Elem().Interface()
			fmt.Println(v)
			switch v.(type) {
			case string:
				fmt.Println("----------------命中了-----------")
				return v.(string), nil
			case int, int8, int32, int64:
				return fmt.Sprintf("%d", v.(int)), nil
			}
		}
		/*fmt.Println(elem)
		switch elem.Kind() {
		case reflect.String:
			return elem.String(), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(elem.Int(), 10), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(elem.Uint(), 10), nil
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(elem.Float(), 'f', -1, 64), nil
		default:
			msg := "no support ptr type " + field.Kind().String() + ", " + elem.Kind().String()
			return "", errors.New(msg)
		}
		*/
	}
	return "", nil
}

func StrInArr(targetStr string, fromArr []string) bool {
	for _, field := range fromArr {
		if field == targetStr {
			return true
		}
	}
	return false
}
