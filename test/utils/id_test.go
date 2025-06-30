package utils

import (
	"fmt"
	"github.com/ihangsen/common/src/utils/id"
	"testing"
)

func TestId(f *testing.T) {
	merge := id.Merge(1, 1)
	fmt.Println(merge)
	index, i := id.Split(merge)
	fmt.Println(index, i)
	code := id.ToCode(merge)
	fmt.Println(code)
	fmt.Println(id.ToId(code))
}
