package game

import (
	"math/rand"
)

type Entity struct {
	Field  [][]int
	Target int
	Width  int
	Score int
}

func (g *Entity) InitFields() {
	field := make([][]int, g.Width)
	for i := 0; i < g.Width; i++ {
		field[i] = make([]int, g.Width)
	}
	g.Field = field
	g.spawnValue()
	g.spawnValue()
}

func (g *Entity) Reset() {
	g.InitFields()
	g.Score = 0
}

func (g *Entity) blankFields() (ret [][2]int) {
	for i := 0; i < g.Width; i++ {
		for j := 0; j < g.Width; j++ {
			if g.Field[i][j] == 0 {
				ret = append(ret, [2]int{i, j})
			}
		}
	}
	return
}

func (g *Entity) spawnValue() {
	newVal := 2
	// 1/9 probability to spawnValue 4
	if rand.Intn(100) > 90 {
		newVal = 4
	}
	blanks := g.blankFields()
	if len(blanks) == 0 {
		return
	}
	blank := blanks[random(0, len(blanks))]
	row, col := blank[0], blank[1]
	g.Field[row][col] = newVal
}

func (g *Entity) addScore(v int) {
	g.Score += v
}

func moveLeft(field [][]int, g *Entity) [][]int {
	mergeRow := func(row []int) {
		for i, v := range row {
			if i+1 >= len(row) {
				return
			}
			if row[i] == 0 {
				row[i] = row[i+1]
				row[i+1] = 0
			}
			if row[i+1] == v {
				row[i] += row[i+1]
				row[i+1] = 0
				if g != nil {
					g.addScore(v << 1)
				}
			}
		}
	}
	compressedField := compress(field)
	for _, row := range compressedField {
		mergeRow(row)
	}
	return compressedField
}

func moveRight(field [][]int, g *Entity) [][]int {
	return invert(moveLeft(invert(field), g))
}

func (g *Entity) MoveLeft() {
	if g.GameOver() {
		return
	}
	newFiled := moveLeft(g.Field, g)
	g.Field = newFiled
	g.spawnValue()
}

func (g *Entity) MoveRight() {
	if g.GameOver() {
		return
	}
	newField := moveRight(g.Field, g)
	g.Field = newField
	g.spawnValue()
}

func (g *Entity) MoveUp() {
	if g.GameOver() {
		return
	}
	newField := transpose(moveLeft(transpose(g.Field), g))
	g.Field = newField
	g.spawnValue()
}

func (g *Entity) MoveDown() {
	if g.GameOver() {
		return
	}
	newField := transpose(moveRight(transpose(g.Field), g))
	g.Field = newField
	g.spawnValue()
}

func (g *Entity) GameOver() bool {
	if g.Win() {
		return false
	}
	if len(g.blankFields()) > 0 {
		return false
	}
	for i := 0; i < g.Width; i++ {
		for j := 0; j < g.Width; j++ {
			if (i-1 >= 0 && g.Field[i][j] == g.Field[i-1][j]) ||
				(i+1 < g.Width && g.Field[i][j] == g.Field[i+1][j]) ||
				(j+1 < g.Width && g.Field[i][j] == g.Field[i][j+1]) {
				return false
			}
		}
	}
	return true
}

func (g *Entity) Win() bool {
	for _, row := range g.Field {
		for _, v := range row {
			if v == g.Target {
				return true
			}
		}
	}
	return false
}
