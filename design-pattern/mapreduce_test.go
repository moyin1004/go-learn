package design

import (
	"strings"
	"testing"
)

var strs []string = []string{"Hao", "Chen", "MegaEase", "Hao",
	"fjaldjfklsajklahsgjkhdsjfhajkhfjkasdlhfbglhaghsklaheDDDFdajgksadlfdjs",
	"Chen", "MegaEase", "Hao", "Chen", "MegaEase"}

func BenchmarkMap(b *testing.B) {
	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	Map(strs, upcase)
}

func BenchmarkMapByPtr(b *testing.B) {
	upcase := func(s *string) string {
		return strings.ToUpper(*s)
	}
	MapByPtr(strs, upcase)
}
