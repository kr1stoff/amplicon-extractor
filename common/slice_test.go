package common

import (
	"fmt"
	"os"
	"testing"
)

func TestRemoveDupplicates(t *testing.T) {
	slice := []int{5533, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5534, 5535}
	// slice := []int{}
	uniqueSlice := RemoveDupplicates(slice)
	fmt.Fprintln(os.Stdout, uniqueSlice)
	// 判断 slice 为否为空
	if len(uniqueSlice) == 0 {
		t.Errorf("Expected non-empty slice, got empty slice")
	}
}
