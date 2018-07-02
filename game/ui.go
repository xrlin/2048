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
	entity.InitFields()
	stage.entity = entity
}

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


func (stage *Stage) GameOver() bool {
	return stage.entity.GameOver()
}

func (stage *Stage) FieldWidth() int {
	return stage.entity.Width
}

func (stage *Stage) fieldWidget() *ui.Row {
	table := ui.NewTable()
	table.Rows = stage.Rows()
	table.FgColor = ui.ColorWhite
	table.BgColor = ui.ColorDefault
	table.TextAlign = ui.AlignCenter
	table.Separator = false
	table.Analysis()
	table.SetSize()
	table.Border = false
	return ui.NewCol(2, 0, table)
}

func (stage *Stage) Status() string {
	ret := "Playing..."
	if stage.GameOver() {
		ret = "GameOver"
	}
	return ret
}

func (stage *Stage) infoWidget() *ui.Row {
	info := "Score: " +  strconv.Itoa(stage.entity.Score) + "\n\n\n"+  stage.Status() + "\n\n\n\nUse w, a, s, d or\n left, right, up, down arrow to control;\n r to reset[fg-red]"
	par := ui.NewPar(info)
	par.BorderLabel = "Info"
	par.Border = false
	par.Height = 20
	par.BorderFg = ui.ColorYellow

	return ui.NewCol(6, 0, par)
}

func (stage *Stage) StatusWidget() *ui.Row {
	info := "Score: " +  strconv.Itoa(stage.entity.Score) + "\n\n\n"+ "Playing...\n\n\n\nUse w, a, s, d or left, right, up, down arrow to control; r to reset[fg-red]"
	par := ui.NewPar(info)
	par.BorderLabel = "Status"
	par.Height = 10
	par.BorderFg = ui.ColorBlue
	return ui.NewCol(2, 0, par)
}

func (stage *Stage) draw()  {
	// avoid deprecated render
	ui.Body.Rows = make([]*ui.Row, 0)

	ui.Body.AddRows(ui.NewRow(stage.fieldWidget(), stage.infoWidget()))
	ui.Body.Align()
	ui.Render(ui.Body)
}

func (stage *Stage) Reset() {
	logger.Println("resetint...")
	stage.entity.Reset()
	stage.draw()
}

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

func (stage *Stage) Rows() [][]string {
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

