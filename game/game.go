package game

import (
	"math/rand"
	"time"
)

// Game manages the overall game state and logic.
type Game struct {
	Board *Board
	Snake *Snake
	Food  Point
	State GameState
	score int
	rng   *rand.Rand // Random number generator for food placement
}

// Config holds the configuration for creating a new game.
type Config struct {
	BoardWidth    int
	BoardHeight   int
	InitialLength int
}

// NewGame initializes a new game instance.
func NewGame(cfg Config) *Game {
	board := NewBoard(cfg.BoardWidth, cfg.BoardHeight)
	// Start snake in the middle, moving right
	startPoint := Point{X: cfg.BoardWidth / 2, Y: cfg.BoardHeight / 2}
	snake := NewSnake(startPoint, cfg.InitialLength, Right)

	// Seed the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	game := &Game{
		Board: board,
		Snake: snake,
		State: Running,
		score: 0,
		rng:   rng,
	}
	game.spawnFood() // Place initial food
	return game
}

// spawnFood places food at a random location not occupied by the snake.
func (g *Game) spawnFood() {
	var p Point
	for {
		p = Point{
			X: g.rng.Intn(g.Board.Width),
			Y: g.rng.Intn(g.Board.Height),
		}
		// Ensure food doesn't spawn on the snake
		if !g.Snake.IsOnSnake(p) {
			g.Food = p
			return
		}
	}
}

// Tick advances the game state by one step.
// It handles snake movement, collisions, and food consumption.
func (g *Game) Tick() {
	if g.State != Running {
		return // Game is already over
	}

	// Move the snake
	nextHead := g.Snake.Move()

	// Check for wall collisions
	if g.Board.IsOutOfBounds(nextHead) {
		g.State = Lost
		return
	}

	// Check for self-collisions
	if g.Snake.CheckSelfCollision() {
		g.State = Lost
		return
	}

	// Check for food consumption
	if nextHead == g.Food {
		g.score++
		g.Snake.Grow(1) // Grow by one segment
		g.spawnFood()   // Spawn new food
	}
}

// SetDirection sets the snake's direction.
func (g *Game) SetDirection(dir Direction) {
	if g.State == Running {
		g.Snake.SetDirection(dir)
	}
}

// GetScore returns the current score.
func (g *Game) GetScore() int {
	return g.score
}

// GetState returns the current game state.
func (g *Game) GetState() GameState {
	return g.State
}
