package string

import (
	"strings"
	"testing"
)

var s = strings.Repeat("a", 1024)

func test() {
	b := []byte(s)
	_ = string(b)
}

func test2() {
	b := Str2bytes(s)
	_ = Bytes2str(b)
}

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}

// go test -run=xxx -bench=. -benchtime="3s" -cpuprofile profile_cpu.out
// go test -bench=.

/*************

goos: windows
goarch: amd64
pkg: github.com/gogoods/x/string
BenchmarkTest1-8          5000000               277 ns/op
BenchmarkTest2-8    1000000000               2.82 ns/op
PASS
ok      github.com/gogoods/x/string     4.819s


***************/
