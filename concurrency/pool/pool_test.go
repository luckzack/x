package pool

import (
	"fmt"
	"testing"
	"time"
)

func Test_pool(t *testing.T) {
	p := New("test", 3, 1)

	fmt.Println(p.Print())

	try := 5

	for {
		if try > 0 {
			_, err := p.GetConn()
			if err != nil {
				fmt.Println("get conn err:", err.Error())
			} else {
				fmt.Println("get conn succ")
			}
			try--
		} else {
			break
		}
	}

	fmt.Println(p.Print())

	//p.releaseConn()
	time.Sleep(time.Second * 2)

	fmt.Println(p.Print())

	time.Sleep(time.Second * 2)

	fmt.Println(p.Print())
}
