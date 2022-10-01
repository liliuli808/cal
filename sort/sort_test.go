package sort

import (
	"fmt"
	"testing"
)

func TestQuickSOrt(t *testing.T) {
	q := QuickSort([]int{5, 2, 1, 4, 6, 3}, 0, 6)
	fmt.Println(q)
}
