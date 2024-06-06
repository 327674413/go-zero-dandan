package mapd

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
)

func AnyToStruct(mapStringAny any, toStructPt any) error {
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			stringToInt64HookFunc,
		),
		Result:  toStructPt,
		TagName: "mapstructure",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	if err = decoder.Decode(mapStringAny); err != nil {
		return err
	}
	return nil
}
func stringToInt64HookFunc(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String && t.Kind() == reflect.Int64 {
		return strconv.ParseInt(data.(string), 10, 64)
	}
	return data, nil
}
