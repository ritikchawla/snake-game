package game

// Point represents a coordinate on the game board.
type Point struct {
	X int
	Y int
}

// Direction represents the snake's movement direction.
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// String returns a string representation of the GameState.
func (s GameState) String() string {
	switch s {
	case Running:
		return "Running"
	case Lost:
		return "Lost"
	case Won:
		return "Won"
	default:
		return "Unknown"
	}
}

// GameState represents the current state of the game.
type GameState int

const (
	Running GameState = iota
	Lost
	Won // Optional, maybe if the snake fills the board
)
