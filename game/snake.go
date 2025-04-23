package game

import "container/list"

// Snake represents the player-controlled snake.
type Snake struct {
	Body      *list.List // Using a linked list for efficient addition/removal at both ends
	Direction Direction
	grow      int // Counter for how many segments to grow
}

// NewSnake creates a new snake at a starting position.
func NewSnake(start Point, initialLength int, initialDirection Direction) *Snake {
	body := list.New()
	// Add the head first
	body.PushFront(start)
	// Add the rest of the body segments behind the head, placed to its left
	for i := 1; i < initialLength; i++ {
		body.PushBack(Point{X: start.X - i, Y: start.Y})
	}
	return &Snake{
		Body:      body,
		Direction: initialDirection,
		grow:      0,
	}
}

// Head returns the position of the snake's head.
func (s *Snake) Head() Point {
	return s.Body.Front().Value.(Point)
}

// Move advances the snake one step in its current direction.
// It returns the new head position.
func (s *Snake) Move() Point {
	head := s.Head()
	var next Point

	switch s.Direction {
	case Up:
		next = Point{X: head.X, Y: head.Y - 1}
	case Down:
		next = Point{X: head.X, Y: head.Y + 1}
	case Left:
		next = Point{X: head.X - 1, Y: head.Y}
	case Right:
		next = Point{X: head.X + 1, Y: head.Y}
	}

	// Add new head
	s.Body.PushFront(next)

	// Handle growth or remove tail
	if s.grow > 0 {
		s.grow--
	} else {
		s.Body.Remove(s.Body.Back())
	}

	return next
}

// SetDirection changes the snake's direction, preventing direct reversal.
func (s *Snake) SetDirection(newDir Direction) {
	currentDir := s.Direction
	// Prevent reversing direction
	if (newDir == Up && currentDir == Down) ||
		(newDir == Down && currentDir == Up) ||
		(newDir == Left && currentDir == Right) ||
		(newDir == Right && currentDir == Left) {
		return
	}
	s.Direction = newDir
}

// Grow increases the snake's length on the next move.
func (s *Snake) Grow(amount int) {
	s.grow += amount
}

// CheckSelfCollision checks if the snake's head has collided with its body.
func (s *Snake) CheckSelfCollision() bool {
	head := s.Head()
	element := s.Body.Front().Next() // Start checking from the second segment
	for element != nil {
		if element.Value.(Point) == head {
			return true
		}
		element = element.Next()
	}
	return false
}

// IsOnSnake checks if a given point is part of the snake's body.
func (s *Snake) IsOnSnake(p Point) bool {
	for element := s.Body.Front(); element != nil; element = element.Next() {
		if element.Value.(Point) == p {
			return true
		}
	}
	return false
}
