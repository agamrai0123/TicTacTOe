package main

import (
	"image/color"
	"strconv"

	"github.com/go-p5/p5"
)

func main() {
	p5.Run(setup, draw)
}

func setup() {
	p5.Canvas(600, 600)
	p5.Background(color.Gray{Y: 220})

}

func draw() {
	coord := [9]string{}
	state := "x"
	drawGrid()

	switch {
	case p5.Event.Mouse.Pressed:
		N := getGrid(p5.Event.Mouse.Position.X, p5.Event.Mouse.Position.Y)
		p5.Text(strconv.Itoa(N), 300, 300)
		x, y := getCoord(N)
		p5.Text(strconv.Itoa(N), x, y)
		coord[N-1] = state
		if state == "x" {
			state = "o"
		} else {
			state = "x"
		}
	default:
	}

	states(coord[:])
	// Win(1, 9)
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

func getCoord(n int) (float64, float64) {
	gridSize := 200
	row := (n - 1) / 3
	col := (n - 1) % 3
	return float64(col)*float64(gridSize) + float64(gridSize)/2, float64(row)*float64(gridSize) + float64(gridSize)/2
}

func drawCircle(x float64, y float64) {
	p5.StrokeWidth(4)
	p5.Stroke(color.Black)
	p5.Circle(x, y, 100)
}

func drawCross(x float64, y float64) {
	p5.StrokeWidth(4)
	p5.Stroke(color.Black)
	p5.Line(x-50, y-50, x+50, y+50)
	p5.Line(x-50, y+50, x+50, y-50)
}

func drawGrid() {
	p5.StrokeWidth(4)
	p5.Stroke(color.Black)
	p5.Line(0, 200, 600, 200)
	p5.Line(0, 400, 600, 400)
	p5.Line(200, 0, 200, 600)
	p5.Line(400, 0, 400, 600)
}

func states(c []string) {
	// for i := 1; i <= 9; i++ {
	// 	x, y := getCoord(i)
	// 	if i%2 == 0 {
	// 		drawCircle(x, y)
	// 	} else {
	// 		drawCross(x, y)
	// 	}
	// }
	for i := 0; i < 9; i++ {
		if c[i] == "x" {
			drawCross(getCoord(i + 1))
		} else if c[i] == "o" {
			drawCircle(getCoord(i + 1))
		}
	}
}

func Win(p1 int, p2 int) {
	p5.StrokeWidth(30)
	p5.Stroke(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	x1, y1 := getCoord(p1)
	x2, y2 := getCoord(p2)

	p5.Line(x1, y1, x2, y2)
}
