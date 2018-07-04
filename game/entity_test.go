package game

import (
	"testing"
	"reflect"
)

func TestEntity_moveLeft(t *testing.T) {
	testCases := []struct {
		Input    [][]int
		Expected [][]int
	}{
		{[][]int{
			{0, 0, 2, 4},
			{2, 0, 2, 2},
			{4, 4, 4, 4},
			{8, 4, 4, 0},
		},
			[][]int{
				{2, 4, 0, 0},
				{4, 2, 0, 0},
				{8, 8, 0, 0},
				{8, 8, 0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		if out := moveLeft(testCase.Input, nil); !reflect.DeepEqual(out, testCase.Expected) {
			t.Errorf("Input with %v. Expected: %v, but have %v", testCase.Input, testCase.Expected, out)
		}
	}
}