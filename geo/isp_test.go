package geo

import (
	"fmt"
	"testing"
)

func Test_xx(t *testing.T) {
	Init("./../../../var/mydata4vipday2-4.dat")
	fmt.Println(GetLocation("115.159.66.125"))
}
