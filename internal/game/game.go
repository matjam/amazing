package game

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/matjam/amazing/internal/life"
)

type Game struct {
	width, height int

	board *life.Board

	mouseDown bool
	pause     bool
}

func NewGame(width int, height int) *Game {
	g := new(Game)
	g.width = width
	g.height = height
	g.board = life.NewLife(width, height)
	return g
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.mouseDown = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.mouseDown = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		g.board.Seed(10) // 10%
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.pause = !g.pause
	}

	if !g.pause {
		g.board.Step()
	}

	if g.mouseDown {
		x, y := ebiten.CursorPosition()
		for yy := y - 2; yy < y+3; yy++ {
			for xx := x - 2; xx < x+3; xx++ {
				g.board.SetCell(xx, yy, life.CellTypeAlive)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var y, x int

	for y = 0; y < g.board.Height; y++ {
		for x = 0; x < g.board.Width; x++ {
			if g.board.GetCell(x, y) == life.CellTypeAlive {
				screen.Set(int(x), int(y), color.RGBA{G: 0xff, A: 0xff})
			} else {
				screen.Set(int(10), int(10), color.RGBA{A: 0xff})
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.width), int(g.height)
}
