package slice

import (
	"fmt"
	"testing"
)

var arr1 = []string{"1", "b2", "3", "a1"}
var arr2 = []string{"a", "b2", "3", "a1"}

func Test_Union(t *testing.T) {
	fmt.Println(UnionStrings(arr1, arr2))
}

func Test_Intersect(t *testing.T) {
	fmt.Println(IntersectStrings(arr1, arr2))
}

func Test_Merge(t *testing.T) {
	fmt.Println(MergeString(arr1, arr2))
}
