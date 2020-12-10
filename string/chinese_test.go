package utils

import (
	"testing"
)

var sample1 = "This is China,这是中国"
var sample2 = "1哈哈哈哈"
var sample3 = "=，1x c d"
var sample4 = "哈11111"
var sample5 = "11ie üe er"

// go test -bench=. -benchmem

func Benchmark_ContainsChineseByUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsChineseByUnicode(sample1)
	}
}
func Benchmark_ContainsChineseByRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsChineseByRegexp(sample1)
	}
}

// go test ./ -test.run=Test_ContainsChineseByUnicode -v
func Test_ContainsChineseByUnicode(t *testing.T) {
	t.Log(containsChineseByUnicode(sample1))
	t.Log(containsChineseByUnicode(sample2))
	t.Log(containsChineseByUnicode(sample3))
	t.Log(containsChineseByUnicode(sample4))
	t.Log(containsChineseByUnicode(sample5))
}

// go test ./ -test.run=Test_ContainsChineseByRegexp -v
func Test_ContainsChineseByRegexp(t *testing.T) {
	t.Log(containsChineseByRegexp(sample1))
	t.Log(containsChineseByRegexp(sample2))
	t.Log(containsChineseByRegexp(sample3))
	t.Log(containsChineseByRegexp(sample4))
	t.Log(containsChineseByRegexp(sample5))
}
