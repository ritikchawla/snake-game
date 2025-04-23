# Snake Game (Go + React)

A classic Snake game implemented with a Go backend and a React frontend using Vite. Communication between the frontend and backend is handled via WebSockets.

## Features

*   Classic Snake gameplay.
*   Real-time updates using WebSockets.
*   Go backend managing game logic (board, snake, food, collisions, state).
*   React frontend for rendering the game and handling user input.
*   High score tracking using browser `localStorage`.

## Technologies Used

*   **Backend:** Go
    *   `gorilla/websocket` for WebSocket communication.
*   **Frontend:** React (with Vite)
    *   JavaScript (ES6+)
    *   CSS

## Project Structure

```
.
├── .gitignore        # Git ignore rules
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── main.go           # Go backend server entrypoint
├── game/             # Go package for core game logic
│   ├── board.go
│   ├── game.go
│   ├── models.go
│   └── snake.go
├── frontend/         # React frontend application (Vite)
│   ├── public/
│   ├── src/
│   │   ├── components/
│   │   │   ├── GameBoard.css
│   │   │   └── GameBoard.jsx
│   │   ├── App.css
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── index.html
│   ├── package.json
│   └── vite.config.js
└── README.md         # This file
```

## Getting Started

### Prerequisites

*   Go (version 1.18 or later recommended)
*   Node.js and npm (or yarn)

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/ritikchawla/snake-game
    cd snake-game
    ```

2.  **Backend Setup:**
    *   Navigate to the project root directory (e.g., `snake-game`).
    *   Tidy dependencies:
        ```bash
        go mod tidy
        ```
    *   Run the backend server:
        ```bash
        go run main.go
        ```
        The server will start on `http://localhost:8080`.

3.  **Frontend Setup:**
    *   Open a **new terminal window**.
    *   Navigate to the `frontend` directory:
        ```bash
        cd frontend
        ```
    *   Install dependencies:
        ```bash
        npm install
        ```
    *   Run the frontend development server:
        ```bash
        npm run dev
        ```
        This will usually open the game automatically in your browser at `http://localhost:5173` (or similar).

4.  **Play:**
    *   Open your browser to the frontend URL (e.g., `http://localhost:5173`).
    *   Use Arrow Keys or WASD keys to control the snake.
    *   Press 'R' to restart the game after "Game Over" or to attempt reconnection if disconnected.
