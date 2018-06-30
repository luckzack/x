package weight

import (
	"testing"
	"fmt"

)

func Test_x(t *testing.T){

	weights := []int{2,4,5}

	c1 := 0
	c2 := 0
	c3 := 0

	for i:=0;i<100000;i++ {
		v := SelectOne(weights...)
		if v == 0 {
			c1++
		}else if v == 1 {
			c2++
		}else if v == 2 {
			c3++
		}
	}

	fmt.Println(c1, c2, c3)

}

func Test_y(t *testing.T){
	c, err := NewCalculator([][]int{[]int{0,9,50}, []int{10,19,30}, []int{20,39,20}})

	if err != nil {
		fmt.Println(err)
		return
	}

	c1 := 0
	c2 := 0
	c3 := 0
	for i:=0;i<10;i++ {
		idx, v := c.Calculate()
		fmt.Println(v)
		if idx == 0 {
			c1++
		}else if idx == 1 {
			c2++
		}else if idx == 2 {
			c3++
		}
	}

	fmt.Println("->", c1, c2, c3)

}



