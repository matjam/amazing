package grid

import "github.com/rs/zerolog/log"

type CellType uint8

type Grid struct {
	allowedCellTypes []bool
	cells            [][]CellType
	Width            int
	Height           int
}

func NewGrid(width int, height int, allowedCellTypes []CellType) *Grid {
	grid := new(Grid)
	grid.Width = width
	grid.Height = height

	grid.cells = make([][]CellType, grid.Height)
	for y := range grid.cells {
		grid.cells[y] = make([]CellType, grid.Width)
	}

	grid.allowedCellTypes = make([]bool, 256)
	for cellType := range allowedCellTypes {
		grid.allowedCellTypes[cellType] = true
	}
	return grid
}

// SetCell sets the CellType of the cell at the given location.
func (g *Grid) SetCell(x int, y int, cellType CellType) {
	if x >= g.Width || y >= g.Height || x < 0 || y < 0 {
		log.Warn().
			Int("x", x).
			Int("y", y).
			Uint8("cellType", uint8(cellType)).
			Msg("attempt to set cell out of bounds")
		return
	}

	if !g.allowedCellTypes[cellType] {
		log.Warn().
			Int("x", x).
			Int("y", y).
			Uint8("cellType", uint8(cellType)).
			Msg("attempt to set illegal cell type")
		return
	}

	g.cells[y][x] = cellType
}

// GetCell will return the CellType of the given location.
func (g *Grid) GetCell(x int, y int) CellType {
	if x >= g.Width || y >= g.Height || x < 0 || y < 0 {
		log.Warn().
			Int("x", x).
			Int("y", y).
			Msg("attempt to get cell out of bounds")
		return 0
	}
	return g.cells[y][x]
}

// Clear the entire grid - setting to whatever CellType 0 is.
func (g *Grid) Clear() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.SetCell(x, y, 0)
		}
	}
}
