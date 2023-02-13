package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Point is a struct for representing 2D vectors.
type Point struct {
	x, y float64
}

type Line struct {
	pos1, pos2 Point
	color      color.RGBA
}

// NewLine initializes and returns a new Line instance.
func NewLine(x1, y1, x2, y2 float64) *Line {
	return &Line{
		pos1: Point{x: x1, y: y1},
		pos2: Point{x: x2, y: y2},
		color: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0xff,
			A: 0xff,
		},
	}
}

func Draw(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	ebitenutil.DrawLine(img, x1, y1, x2, y2, c)
}

func (l *Line) Update() {
	newx := l.pos2.x*math.Cos(math.Pi/180) - l.pos2.y*math.Sin(math.Pi/180)
	newy := l.pos2.x*math.Sin(math.Pi/180) + l.pos2.y*math.Cos(math.Pi/180)
	l.pos2.x, l.pos2.y = newx, newy
}

func (l *Line) Draw(screen *ebiten.Image) {
	Draw(screen, l.pos1.x, l.pos1.y, l.pos2.x, l.pos2.y, l.color)
}

// Game is a game instance.
type Game struct {
	width, height int
	line          *Line
}

// NewGame returns a new Game instance.
func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		line:   NewLine(float64(width/2), float64(height/2), 430, 130),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

// Update updates a game state.
func (g *Game) Update() error {
	g.line.Update()
	return nil
}

// Draw renders a game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.line.Draw(screen)
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
