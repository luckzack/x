package workflow

import (
	"fmt"
	"testing"
)

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func baz() {
	fmt.Println("baz")
}

func Test_xx(t *testing.T) {

	Once(foo)
	Once(foo)
	Once(bar)
	Once(bar)
	Once(foo)
	Once(bar)
	Once(bar)
	Once(baz)
	Once(baz)
}
