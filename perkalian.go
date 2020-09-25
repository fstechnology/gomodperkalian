package gomodperkalian

import "reflect"

func Perkalian(items ...int64) int64 {
	res := int64(1)
	for _, item := range items {
		res *= item
	}
	return res
}

func PerkalianInt(items ...int) int64 {
	res := 1
	for _, item := range items {
		res *= item
	}
	return int64(res)
}

func convertToInt64(data interface{}) int64 {
	if data == nil {
		return 0
	}
	int64Type := reflect.TypeOf(int64(0))
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(int64Type) {
		return 0
	}
	res := v.Convert(int64Type)
	return res.Int()
}
