package game

import (
	"testing"
	"reflect"
)

func TestCompress(t *testing.T) {
	testCases := []struct {
		Input    [][]int
		Expected [][]int
	}{
		{[][]int{
			{0, 0, 2, 4},
			{0, 0, 2, 8},
			{0, 0, 4, 4},
			{2, 2, 2, 4},
		},
			[][]int{
				{2, 4, 0, 0},
				{2, 8, 0, 0},
				{4, 4, 0, 0},
				{2, 2, 2, 4},
			},
		},
		{[][]int{
			{0, 2, 0, 4},
			{2, 0, 0, 8},
			{4, 0, 4, 0},
			{2, 4, 0, 0},
		},
			[][]int{
				{2, 4, 0, 0},
				{2, 8, 0, 0},
				{4, 4, 0, 0},
				{2, 4, 0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		if out := Compress(testCase.Input); !reflect.DeepEqual(out, testCase.Expected) {
			t.Errorf("Input with %v. Expected: %v, but have %v", testCase.Input, testCase.Expected, out)
		}
	}
}

func TestInvert(t *testing.T) {
	testCases := []struct {
		Input    [][]int
		Expected [][]int
	}{
		{[][]int{
			{0, 0, 2, 4},
			{0, 8, 2, 8},
			{4, 0, 0, 2},
			{0, 0, 2, 4},
		},
			[][]int{
				{4, 2, 0, 0},
				{8, 2, 8, 0},
				{2, 0, 0, 4},
				{4, 2, 0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		if out := Invert(testCase.Input); !reflect.DeepEqual(out, testCase.Expected) {
			t.Errorf("Input with %v. Expected: %v, but have %v", testCase.Input, testCase.Expected, out)
		}
	}
}

func TestTranspose(t *testing.T) {
	testCases := []struct {
		Input    [][]int
		Expected [][]int
	}{
		{[][]int{
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
			{1, 1, 1, 1},
		},
			[][]int{
				{1, 1, 1, 1},
				{0, 1, 1, 1},
				{0, 0, 1, 1},
				{0, 0, 0, 1},
			},
		},
		{[][]int{
			{0, 0, 0, 8},
			{1, 0, 0, 0},
			{1, 2, 0, 0},
			{1, 1, 1, 0},
		},
			[][]int{
				{0, 1, 1, 1},
				{0, 0, 2, 1},
				{0, 0, 0, 1},
				{8, 0, 0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		if out := Transpose(testCase.Input); !reflect.DeepEqual(out, testCase.Expected) {
			t.Errorf("Input with %v. Expected: %v, but have %v", testCase.Input, testCase.Expected, out)
		}
	}
}