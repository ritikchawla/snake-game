package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ritikchawla/snake-game/game" // Updated import path

	"github.com/gorilla/websocket"
)

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity in development
		// In production, restrict this to your frontend's origin
		return true
	},
}

// Client represents a single connected WebSocket client and their game state.
type Client struct {
	conn *websocket.Conn
	game *game.Game
	mu   sync.Mutex // To protect concurrent writes to the connection
}

// GameStatePayload is the structure sent to the client over WebSocket.
type GameStatePayload struct {
	BoardWidth  int          `json:"boardWidth"`
	BoardHeight int          `json:"boardHeight"`
	SnakeBody   []game.Point `json:"snakeBody"`
	Food        game.Point   `json:"food"`
	Score       int          `json:"score"`
	GameState   string       `json:"gameState"` // Send string representation
}

// handleConnections upgrades HTTP requests to WebSocket connections and starts the game loop.
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer ws.Close()

	log.Println("Client connected")

	// Game Configuration (could be made configurable)
	config := game.Config{
		BoardWidth:    30, // Larger board for better gameplay
		BoardHeight:   20,
		InitialLength: 3,
	}
	g := game.NewGame(config)

	client := &Client{
		conn: ws,
		game: g,
	}

	// Goroutine to handle incoming messages (player input)
	go client.handleInput()

	// Start the game loop for this client
	client.gameLoop()

	log.Println("Client disconnected")
}

// gameLoop runs the game logic and sends state updates to the client.
func (c *Client) gameLoop() {
	// Adjust tick speed for gameplay (faster tick = smoother perceived movement)
	ticker := time.NewTicker(120 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		c.game.Tick() // Advance game state

		// Prepare payload
		gameState := c.game.GetState()
		snakeList := c.game.Snake.Body
		snakeBodyPoints := make([]game.Point, 0, snakeList.Len())
		for e := snakeList.Front(); e != nil; e = e.Next() {
			snakeBodyPoints = append(snakeBodyPoints, e.Value.(game.Point))
		}

		payload := GameStatePayload{
			BoardWidth:  c.game.Board.Width,
			BoardHeight: c.game.Board.Height,
			SnakeBody:   snakeBodyPoints,
			Food:        c.game.Food,
			Score:       c.game.GetScore(),
			GameState:   gameState.String(), // Use the String() method
		}

		// Send state to client
		c.mu.Lock()
		err := c.conn.WriteJSON(payload)
		c.mu.Unlock()

		if err != nil {
			log.Printf("Error sending game state: %v", err)
			break // Stop loop on error (client likely disconnected)
		}

		// Stop loop if game is over
		if gameState != game.Running {
			log.Printf("Game Over for client. Score: %d", c.game.GetScore())
			// Optionally send a final "Game Over" message before breaking
			// c.conn.WriteJSON(...)
			break
		}
	}
}

// handleInput reads messages from the client (direction changes).
func (c *Client) handleInput() {
	defer func() {
		// Ensure connection is closed if this goroutine exits
		c.conn.Close()
	}()

	for {
		// Read message from client
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Client read error: %v", err)
			}
			break // Exit loop on error or close
		}

		// Process message (expecting simple direction strings)
		direction := string(message)
		// log.Printf("Received direction: %s", direction) // Debugging
		switch direction {
		case "UP":
			c.game.SetDirection(game.Up)
		case "DOWN":
			c.game.SetDirection(game.Down)
		case "LEFT":
			c.game.SetDirection(game.Left)
		case "RIGHT":
			c.game.SetDirection(game.Right)
		default:
			log.Printf("Received unknown command: %s", direction)
		}
	}
}

func main() {
	log.Println("Starting Snake Game WebSocket Server on :8080")

	// Serve static files from the frontend build directory (optional, for deployment)
	// fs := http.FileServer(http.Dir("./frontend/dist"))
	// http.Handle("/", fs)

	// WebSocket endpoint
	http.HandleFunc("/ws", handleConnections)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
