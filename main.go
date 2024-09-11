package main

import (
	"image/color"
	"os"
	"time"

	"github.com/go-p5/p5"
)

var (
	coord       [9]string
	state       = "x"
	prevPressed = false
	gameOver    = false
	winPattern  [3]int
)

func main() {
	p5.Run(setup, draw)
}

func setup() {
	p5.Canvas(600, 600)
	p5.Background(color.Gray{Y: 220})
}

func draw() {
	if gameOver {
		drawGrid()
		states(coord[:])
		drawWinLine(winPattern)
		return
	}

	drawGrid()
	states(coord[:])

	if p5.Event.Mouse.Pressed && !prevPressed {
		mousePressed()
	}
	prevPressed = p5.Event.Mouse.Pressed
}

func mousePressed() {
	N := getGrid(p5.Event.Mouse.Position.X, p5.Event.Mouse.Position.Y)
	if N == 0 || coord[N-1] != "" {
		return
	}

	coord[N-1] = state
	x, y := getCoord(N)
	if state == "x" {
		drawCross(x, y)
		state = "o"
	} else {
		drawCircle(x, y)
		state = "x"
	}

	if winner, pattern := checkWin(); winner != "" {
		println("Winner is: " + winner)
		winPattern = pattern // Store the winning pattern
		drawWinLine(pattern)
		endGame()
		return
	}

	if isDraw() {
		println("The game is a draw!")
		endGame()
	}
}

func checkWin() (string, [3]int) {
	winPatterns := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonals
	}

	for _, pattern := range winPatterns {
		if coord[pattern[0]] != "" &&
			coord[pattern[0]] == coord[pattern[1]] &&
			coord[pattern[1]] == coord[pattern[2]] {
			return coord[pattern[0]], pattern
		}
	}
	return "", [3]int{}
}

func isDraw() bool {
	for _, val := range coord {
		if val == "" {
			return false
		}
	}
	return true
}

func drawWinLine(pattern [3]int) {
	if pattern[0] == -1 {
		return
	}
	x1, y1 := getCoord(pattern[0] + 1)
	x2, y2 := getCoord(pattern[2] + 1)

	p5.StrokeWidth(10)
	p5.Stroke(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	p5.Line(x1, y1, x2, y2)
}

func endGame() {
	gameOver = true
	go func() {
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}()
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
	for i := 0; i < 9; i++ {
		if c[i] == "x" {
			x, y := getCoord(i + 1)
			drawCross(x, y)
		} else if c[i] == "o" {
			x, y := getCoord(i + 1)
			drawCircle(x, y)
		}
	}
}
