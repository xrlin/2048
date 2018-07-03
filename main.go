package main

import (
	"github.com/xrlin/2048/game"
)

func main() {

	gameUI := &game.Stage{}

	gameUI.RegisterEntity(&game.Entity{Width: 4})

	gameUI.Run()
}
