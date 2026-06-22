package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	chars := "Go”Ô—‘±‡≥Ã"
	data := []rune(chars)
	utf8.DecodeRune([]byte(chars))
	fmt.Println(data, len(chars))
	chars2 := string(data[1:])
	fmt.Println(chars2)
	chars3 := strings.Replace(chars, "±‡", "ddd", 1)
	fmt.Println(chars3)
}
