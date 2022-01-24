package weight

import (
	"github.com/gogoods/x/random"
	//"fmt"
)

func SelectOne(weights ...int) int{
	if len(weights) == 1 {
		return 0
	}

	sum := 0
	fields := [][]int{}
	for i, v := range weights{
		sum += v
		if i == 0 {
			fields = append(fields, []int{0, sum})
		}else{
			fields = append(fields, []int{sum-v, sum})
		}

	}

	v1 := random.Int(sum)

	//fmt.Println("===", sum)
	//fmt.Println("fields:", fields)
	//fmt.Println("--->", v1)

	for i, field := range fields {
		if v1 >= field[0] && v1 < field[1]{
			return i
		}
	}

	return 0
}
