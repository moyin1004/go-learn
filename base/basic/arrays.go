package basic

import "fmt"

func printArray(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func TestArrays() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(&arr1)
	printArray(&arr3)
}
