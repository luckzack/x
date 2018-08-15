package string

import "testing"

var num = 1000000
var arr = make([]string, num)

func init() {
	for i := 0; i < num; i++ {
		arr[i] = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	}
}

func BenchmarkTest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatShort(arr...)
	}
}

func BenchmarkTest4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatLong(arr...)
	}
}

// go test -run=xxx -bench=. -benchtime="3s" -cpuprofile profile_cpu.out
// go test -bench=.
