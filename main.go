package main

import (
	"2048/game"
)

func main() {

	gameUI := &game.Stage{}

	gameUI.RegisterEntity(&game.Entity{Width: 4})

	gameUI.Run()
}
