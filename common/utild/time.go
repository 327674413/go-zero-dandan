package utild

//与时间相关的方法
import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetStamp() int64 {
	return time.Now().Unix()
}

// Date 格式化时间,默认当前时间
func Date(format string, timestamp ...int64) string {
	var ts int64
	if len(timestamp) == 0 {
		ts = time.Now().Unix()
	} else {
		ts = timestamp[0]
	}
	t := time.Unix(ts, 0)

	replacer := map[string]string{
		"Y": "2006",
		"y": "06",
		"m": "01",
		"n": "1",
		"d": "02",
		"j": "2",
		"H": "15",
		"h": "03",
		"i": "04",
		"s": "05",
		"w": "0",
	}

	res := ""
	for i := 0; i < len(format); i++ {
		if i < len(format)-1 && format[i:i+2] == "\\n" {
			res += "\n"
			i++
		} else if replace, ok := replacer[format[i:i+1]]; ok {
			res += t.Format(replace)
		} else {
			res += format[i : i+1]
		}
	}

	return res
}

// StrToStamp 字符串转成int64的时间戳
func StrToStamp(dateStr string) int64 {
	// 如果是时间戳，则直接返回
	if timestamp, err := strconv.ParseInt(dateStr, 10, 64); err == nil {
		return timestamp
	}

	// 替换日期字符串中的特殊字符为 "-" 和 ":"
	dateStr = strings.ReplaceAll(dateStr, "/", "-")
	dateStr = strings.ReplaceAll(dateStr, ".", "-")
	dateStr = strings.ReplaceAll(dateStr, " ", "-")
	dateStr = strings.ReplaceAll(dateStr, ":", "-")

	// 尝试使用多种日期格式进行转换
	formats := []string{
		"2006-01-02-15-04-05",
		"2006-01-02-15-04",
		"2006-01-02-15",
		"2006-01-02",
		"2006-01-2-15-04-05",
		"2006-01-2-15-04",
		"2006-01-2-15",
		"2006-1-2-15-04-05",
		"2006-1-2-15-04",
		"2006-1-2-15",
		"2006-1-02-15-04-05",
		"2006-1-02-15-04",
		"2006-1-02-15",
		"2006-1-2",
		"2006-01-2",
		"2006-1-02",
	}

	for _, format := range formats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t.Unix()
		}
	}

	return 0
}

func DiffTimeStr(start, end interface{}) (string, error) {
	s := "秒"
	startTime, err := AnyToTime(start)
	if err != nil {
		return "", nil
	}
	endTime, err := AnyToTime(end)
	if err != nil {
		return "", nil
	}

	if startTime.After(endTime) {
		return "", nil
	}
	if startTime.Equal(endTime) {
		return fmt.Sprintf("0%s", s), nil
	}
	//从这里开始还没优化
	// 判断年月是否超过1年
	if startTime.Year() != endTime.Year() || startTime.Month() != endTime.Month() {
		yearDiff := endTime.Year() - startTime.Year()
		monthDiff := int(endTime.Month()) - int(startTime.Month())

		// 处理月份溢出的情况
		if monthDiff < 0 {
			yearDiff--
			monthDiff += 12
		}

		if yearDiff > 0 && monthDiff > 0 {
			return fmt.Sprintf("%d年%d月%d日 %d时%d分", yearDiff, monthDiff, endTime.Day(), endTime.Hour(), endTime.Minute()), nil
		} else if yearDiff > 0 {
			return fmt.Sprintf("%d年%d日 %d时%d分", yearDiff, endTime.Day(), endTime.Hour(), endTime.Minute()), nil
		} else if monthDiff > 0 {
			return fmt.Sprintf("%d月%d日 %d时%d分", monthDiff, endTime.Day(), endTime.Hour(), endTime.Minute()), nil
		}
	}

	// 计算天数差额
	duration := endTime.Sub(startTime)
	days := int(duration.Hours() / 24)

	// 处理小时和分钟
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%d日 %d时%d分", days, hours, minutes), nil
}

func AnyToTime(anyTime interface{}) (time.Time, error) {
	switch v := anyTime.(type) {
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case int:
		return time.Unix(int64(v), 0), nil
	case time.Time:
		return v, nil
	case string:
		res := StrToStamp(v)
		return time.Unix(res, 0), nil
	default:
		return time.Time{}, fmt.Errorf("时间格式%s错误，仅支持int、int32、int64、Time.time", v)
	}
}

func AnyToInt64(anyV interface{}) int64 {
	switch v := anyV.(type) {
	case string:
		res, err := strconv.Atoi(v)
		if err == nil {
			return int64(res)
		} else {
			return 0
		}
	case int64:
		return v
	case int32:
		return int64(v)
	case int:
		return int64(v)
	case time.Time:
		return v.Unix()
	default:
		return 0
	}
}
func AnyToInt(anyV interface{}) int {
	switch v := anyV.(type) {
	case string:
		res, err := strconv.Atoi(v)
		if err == nil {
			return res
		} else {
			return 0
		}
	case int64:
		return int(v)
	case int32:
		return int(v)
	case int:
		return v
	default:
		return 0
	}
}
