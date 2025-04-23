package game

// Board represents the game grid.
type Board struct {
	Width  int
	Height int
}

// NewBoard creates a new game board.
func NewBoard(width, height int) *Board {
	return &Board{
		Width:  width,
		Height: height,
	}
}

// IsOutOfBounds checks if a point is outside the board boundaries.
func (b *Board) IsOutOfBounds(p Point) bool {
	return p.X < 0 || p.X >= b.Width || p.Y < 0 || p.Y >= b.Height
}
