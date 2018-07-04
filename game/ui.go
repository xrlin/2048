package game

import (
	ui "github.com/gizak/termui"
	"strconv"
)

// Stage is the container to draw ui and receive events to interact with user
type Stage struct {
	entity *Entity
}

// RegisterEntity add the entity for Stage.
// Must call this method first after Stage struct created
func (stage *Stage) RegisterEntity(entity *Entity) {
	entity.InitField()
	stage.entity = entity
}

// Move is a proxy of Entity's move.
// Call corresponding move method with direction.
func (stage *Stage) Move(direction string) {
	switch direction {
	case "left":
		stage.entity.MoveLeft()
	case "right":
		stage.entity.MoveRight()
	case "up":
		stage.entity.MoveUp()
	case "down":
		stage.entity.MoveDown()
	}
	stage.draw()
}


func (stage *Stage) gameOver() bool {
	return stage.entity.GameOver()
}

// FieldWidth returns the filed's width
func (stage *Stage) FieldWidth() int {
	return stage.entity.Width
}

// Display the field
func (stage *Stage) fieldWidget() *ui.Row {
	table := ui.NewTable()
	table.Rows = stage.rows()
	table.FgColor = ui.ColorWhite
	table.BgColor = ui.ColorDefault
	table.TextAlign = ui.AlignCenter
	table.Separator = false
	table.Analysis()
	table.SetSize()
	table.Border = false
	return ui.NewCol(2, 0, table)
}

// Status return the corresponding string representing the status of the game.
func (stage *Stage) Status() string {
	ret := "Playing..."
	if stage.gameOver() {
		ret = "GameOver"
	}
	return ret
}

// Display game info.
func (stage *Stage) infoWidget() *ui.Row {
	info := "Score: " +  strconv.Itoa(stage.entity.Score) + "\n\n\n"+  stage.Status() + "\n\n\n\nUse w, a, s, d or\n left, right, up, down arrow to control;\n r to reset[fg-red]"
	par := ui.NewPar(info)
	par.BorderLabel = "Info"
	par.Border = false
	par.Height = 20
	par.BorderFg = ui.ColorYellow

	return ui.NewCol(6, 0, par)
}

func (stage *Stage) draw()  {
	// avoid deprecated render
	ui.Body.Rows = make([]*ui.Row, 0)

	ui.Body.AddRows(ui.NewRow(stage.fieldWidget(), stage.infoWidget()))
	ui.Body.Align()
	ui.Render(ui.Body)
}

// Reset the game
func (stage *Stage) Reset() {
	logger.Println("resetint...")
	stage.entity.Reset()
	stage.draw()
}

// Listen for keyboard event and call the corresponding procedure.
func (stage *Stage) listen() {
	// handle key q pressing
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd", func(e ui.Event) {
		switch e.Path {
		case "/sys/kbd/<left>", "/sys/kbd/a":
			stage.Move("left")
		case "/sys/kbd/<right>", "/sys/kbd/d":
			stage.Move("right")
		case "/sys/kbd/<up>", "/sys/kbd/w":
			stage.Move("up")
		case "/sys/kbd/<down>", "/sys/kbd/s":
			stage.Move("down")
		case "/sys/kbd/r":
			stage.Reset()
		}
	})
}

func (stage *Stage) rows() [][]string {
	ret := make([][]string, stage.FieldWidth())
	for i, row := range stage.entity.Field {
		strRow := make([]string, len(row))
		for j, v := range row {
			//strRow[j] = "[" + strconv.Itoa(v) + "]" + "(fg-white,bg-black)"
			strRow[j] =  strconv.Itoa(v)
		}
		ret[i] = strRow
	}
	return ret
}

// Run start the game and stays in game's main loop
func (stage *Stage) Run() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()
	stage.listen()
	stage.draw()
	ui.Loop()
}

