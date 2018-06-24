package game

import (
	"time"
	"math/rand"
)

func compress(field [][]int) [][]int {
	ret := make([][]int, len(field))
	for i, v := range field {
		row := make([]int, len(v))
		idx := 0
		for _, v := range v {
			if v == 0 {
				continue
			}
			row[idx] = v
			idx++
		}
		ret[i] = row
	}
	return ret
}

func invert(field [][]int) [][]int {
	rows := len(field)
	if rows == 0 {
		return field

	}
	ret := make([][]int, rows)
	cols := len(field[0])
	for r, row := range field {
		newRow := make([]int, cols)
		for c, v := range row {
			newRow[cols-c-1] = v
		}
		ret[r] = newRow
	}
	return ret
}

func transpose(field [][]int) [][]int {
	rows := len(field)
	cols := len(field[0])
	ret := make([][]int, cols)
	for i := 0; i < cols; i++ {
		row := make([]int, rows)
		for j := 0; j < rows; j ++ {
			row[j] = field[j][i]
		}
		ret[i] = row
	}
	return ret
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}