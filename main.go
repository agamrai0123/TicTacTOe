package main

import (
	"image/color"

	"github.com/go-p5/p5"
)

func main() {
	state := " "
	p5.Run(setup, draw)
}

func setup() {
	p5.Canvas(600, 600)
	p5.Background(color.Gray{Y: 220})
}

func draw() {
	p5.StrokeWidth(4)
	p5.Stroke(color.Black)
	p5.Line(0, 200, 600, 200)
	p5.Line(0, 400, 600, 400)
	p5.Line(200, 0, 200, 600)
	p5.Line(400, 0, 400, 600)
	switch {
	case p5.Event.Mouse.Pressed:
		if p5.Event.Mouse.Buttons.Contain(p5.ButtonLeft) {
			p5.Stroke(color.Black)
			p5.Fill(color.RGBA{R: 255, A: 255})
		}
	default:
		p5.Stroke(nil)
		p5.Fill(color.Transparent)
	}
	p5.Ellipse(
		p5.Event.Mouse.Position.X,
		p5.Event.Mouse.Position.Y,
		80, 80,
	)
	// fmt.Println(getGrid(p5.Event.Mouse.Position.X, p5.Event.Mouse.Position.Y))
	p5.Text(States(), 300, 300)

}

func getGrid(x float64, y float64) int {
	gridSize := 200
	col := int(x) / gridSize
	row := int(y) / gridSize

	if col < 0 || col >= 3 || row < 0 || row >= 3 {
		return 0
	}
	return row*3 + col + 1
}

func States() {
	if state[0] == "X" {
		state[0] = "O"
	} else {
		state = "X"
	}
}
