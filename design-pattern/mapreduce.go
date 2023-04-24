/// @file    mapreduce.go
/// @author  moyin(moyin1004@163.com)
/// @date    2023-04-13 15:59:01

package design

import (
	"fmt"
	"strings"
)

func Map[DataType any](data []DataType, transFunc func(DataType) DataType) []DataType {
	result := make([]DataType, len(data))

	for i := 0; i < len(data); i++ {
		result[i] = transFunc(data[i])
	}
	return result
}

func MapByPtr[DataType any](data []DataType, transFunc func(*DataType) DataType) []DataType {
	result := make([]DataType, len(data))

	for i := 0; i < len(data); i++ {
		result[i] = transFunc(&data[i])
	}
	return result
}

func TestMap() {

	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}

	squared_arr := Map(nums, square)
	fmt.Println(squared_arr)
	//[1 4 9 16]

	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, upcase)
	fmt.Println(upstrs)
	//[HAO CHEN MEGAEASE]
}
