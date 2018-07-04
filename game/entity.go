package game

import (
	"math/rand"
)

// The Entity struct contains the game's field, score etc.
// Realize the game logic.
type Entity struct {
	Field  [][]int
	Target int
	Width  int
	Score int
}

// InitField init the fields for g.
func (g *Entity) InitField() {
	field := make([][]int, g.Width)
	for i := 0; i < g.Width; i++ {
		field[i] = make([]int, g.Width)
	}
	g.Field = field
	g.spawnValue()
	g.spawnValue()
}

// Reset will reset the Field and Score.
func (g *Entity) Reset() {
	g.InitField()
	g.Score = 0
}

// blankFields return the fields with blank(zero) value.
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

// spawnValue generates a value, 2 or 4 in 9 to 1 ratio, and set it in a Random blank field.
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
	blank := blanks[Random(0, len(blanks))]
	row, col := blank[0], blank[1]
	g.Field[row][col] = newVal
}

func (g *Entity) addScore(v int) {
	g.Score += v
}

// moveLeft is a helper method to move left and merge the adjacent value with same value.
// Add up the scores according the values after merged
func moveLeft(field [][]int, g *Entity) [][]int {
	mergeRow := func(row []int) {
		for i, v := range row {
			if i+1 >= len(row) {
				return
			}
			if row[i+1] == v {
				row[i] += row[i+1]
				start := i+1
				// move fields behind from right to left
				for ;start + 1 < len(row); start++{
					row[start] = row[start+1]
				}
				// fill zero value to the last value
				row[start] = 0
				if g != nil {
					g.addScore(v << 1)
				}
			}
		}
	}
	compressedField := Compress(field)
	for _, row := range compressedField {
		mergeRow(row)
	}
	return compressedField
}

// Move right. Just opposite with move left.
func moveRight(field [][]int, g *Entity) [][]int {
	return Invert(moveLeft(Invert(field), g))
}

// MoveLeft merge and move left the field and update the score in g.
func (g *Entity) MoveLeft() {
	if g.GameOver() {
		return
	}
	newFiled := moveLeft(g.Field, g)
	g.Field = newFiled
	g.spawnValue()
}

// MoveRight check and call moveRight then up the field in g.
func (g *Entity) MoveRight() {
	if g.GameOver() {
		return
	}
	newField := moveRight(g.Field, g)
	g.Field = newField
	g.spawnValue()
}

// MoveUp up merge the corresponding fields and update it.
// Transpose and move right on the filed like matrix and finally get the result equal to moveRight.
func (g *Entity) MoveUp() {
	if g.GameOver() {
		return
	}
	newField := Transpose(moveLeft(Transpose(g.Field), g))
	g.Field = newField
	g.spawnValue()
}

// MoveDown down merge the corresponding fields and update it.
// Transpose and move right on the filed like matrix and finally get the result equal to moveRight.
func (g *Entity) MoveDown() {
	if g.GameOver() {
		return
	}
	newField := Transpose(moveRight(Transpose(g.Field), g))
	g.Field = newField
	g.spawnValue()
}

// GameOver checks whether the game is over according whether it is any blank fields to move or adjacent fields can be merged.
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

// Win check is the player reach the target of the game
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
