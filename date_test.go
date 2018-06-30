package utils

import (
	"testing"
	"time"
	"fmt"
)

func Test_DayStartTime(t *testing.T){
	fmt.Println(DayStartTime(time.Now()))

	fmt.Println(DayEndTime(time.Now()))
}
