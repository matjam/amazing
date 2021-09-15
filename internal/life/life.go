package life

import (
	"math/rand"
	"time"

	"github.com/matjam/amazing/internal/grid"
)

// Define what cell types we support for this type of grid.
const (
	CellTypeDead grid.CellType = iota
	CellTypeAlive
)

type Board struct {
	*grid.Grid
	next *grid.Grid

	Generation int
	Population int
}

func NewLife(width int, height int) *Board {
	b := new(Board)
	b.Grid = grid.NewGrid(width, height, []grid.CellType{CellTypeDead, CellTypeAlive})
	b.next = grid.NewGrid(width, height, []grid.CellType{CellTypeDead, CellTypeAlive})

	return b
}

// Seed the board with a random amount of life
func (b *Board) Seed(chance int) {
	var y, x int

	rand.Seed(time.Now().UnixMilli())

	for y = 0; y < b.Height; y++ {
		for x = 0; x < b.Width; x++ {
			if rand.Intn(100) > chance {
				b.SetCell(x, y, CellTypeDead)
			} else {
				b.SetCell(x, y, CellTypeAlive)
			}
		}
	}
}

// Step through to the next generation.
//
//    Any live cell with two or three live neighbours survives.
//    Any dead cell with three live neighbours becomes a live cell.
//    All other live cells die in the next generation. Similarly, all other dead cells stay dead.
//
// The initial pattern constitutes the seed of the system. The first generation is created by applying the above rules
// simultaneously to every cell in the seed, live or dead; births and deaths occur simultaneously, and the discrete
// moment at which this happens is sometimes called a tick. Each generation is a pure function of the preceding one.
// The rules continue to be applied repeatedly to create further generations.
func (b *Board) Step() {
	var y, x int

	for y = 0; y < b.Height; y++ {
		for x = 0; x < b.Width; x++ {
			b.next.SetCell(x, y, b.CellStatus(x, y))
		}
	}

	oldGrid := b.Grid
	b.Grid = b.next
	b.next = oldGrid
}

func (b *Board) CellStatus(x int, y int) grid.CellType {
	liveNeighbourCount := 0

	// the boundary we're checking
	var xLeft, yTop, xRight, yBottom int

	// calculate the sides
	if x == 0 {
		xLeft = b.Width - 1
	} else {
		xLeft = x - 1
	}

	if y == 0 {
		yTop = b.Height - 1
	} else {
		yTop = y - 1
	}

	if x == b.Width-1 {
		xRight = 0
	} else {
		xRight = x + 1
	}

	if y == b.Height-1 {
		yBottom = 0
	} else {
		yBottom = y + 1
	}

	// now count the living
	checkCells := [8]grid.CellType{
		b.GetCell(xLeft, yTop),
		b.GetCell(x, yTop),
		b.GetCell(xRight, yTop),
		b.GetCell(xLeft, y),
		b.GetCell(xRight, y),
		b.GetCell(xLeft, yBottom),
		b.GetCell(x, yBottom),
		b.GetCell(xRight, yBottom),
	}

	for _, v := range checkCells {
		if v == CellTypeAlive {
			liveNeighbourCount++
		}
	}

	// What happens if the cell is live?
	switch b.GetCell(x, y) {
	case CellTypeAlive:
		if liveNeighbourCount > 1 && liveNeighbourCount < 4 {
			return CellTypeAlive
		}
	case CellTypeDead:
		if liveNeighbourCount == 3 {
			return CellTypeAlive
		}
	}

	return CellTypeDead
}
