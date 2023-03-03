package low

import (
	"fmt"
	"testing"
)

func ArrayMax(arr []int) {
	maxValue := arr[0]
	maxValueIndex := 0
	for i := 0; i < len(arr); i++ {
		if maxValue < arr[i] {
			maxValue = arr[i]
			maxValueIndex = i
		}

	}
	fmt.Printf("maxValue=%v maxValueIndex=%v \n", maxValue, maxValueIndex)
}

func TestDay2Main(t *testing.T) {
	array01 := []int{1, 22, 33, 12}
	ArrayMax(array01)
}
