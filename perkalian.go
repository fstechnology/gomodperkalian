package gomodperkalian

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
