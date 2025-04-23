import React, { useState, useEffect, useRef, useCallback } from 'react';
import GameBoard from './components/GameBoard';
import './App.css';

const WEBSOCKET_URL = 'ws://localhost:8080/ws'; // Backend WebSocket URL

function App() {
  const [gameState, setGameState] = useState(null);
  const [isConnected, setIsConnected] = useState(false);
  const [message, setMessage] = useState('Connecting to server...');
  const [highScore, setHighScore] = useState(() => {
    // Load high score from localStorage on initial render
    return parseInt(localStorage.getItem('snakeHighScore') || '0', 10);
  });
  const ws = useRef(null);

  const connectWebSocket = useCallback(() => {
    console.log('Attempting to connect WebSocket...');
    ws.current = new WebSocket(WEBSOCKET_URL);

    ws.current.onopen = () => {
      console.log('WebSocket Connected');
      setIsConnected(true);
      setMessage('Connected! Game starting...');
    };

    ws.current.onclose = () => {
      console.log('WebSocket Disconnected');
      setIsConnected(false);
      setGameState(null); // Clear game state on disconnect
      setMessage('Disconnected. Press R to reconnect.');
      ws.current = null; // Ensure ref is nullified
    };

    ws.current.onerror = (error) => {
      console.error('WebSocket Error:', error);
      setMessage('Connection error. Check if the backend server is running.');
    };

    ws.current.onmessage = (event) => {
      try {
        const receivedState = JSON.parse(event.data);
        setGameState(receivedState);

        // Update high score if game is lost and score is higher
        if (receivedState.gameState === 'Lost') {
          setMessage(`Game Over! Score: ${receivedState.score}. Press R to play again.`);
          if (receivedState.score > highScore) {
            setHighScore(receivedState.score);
            localStorage.setItem('snakeHighScore', receivedState.score.toString());
            setMessage(`Game Over! New High Score: ${receivedState.score}! Press R to play again.`);
          }
        } else if (receivedState.gameState === 'Running') {
          setMessage(`Score: ${receivedState.score}`);
        }
      } catch (error) {
        console.error('Failed to parse game state:', error);
      }
    };
  }, []); // Empty dependency array ensures this function is stable

  // Effect to handle initial connection and cleanup
  useEffect(() => {
    connectWebSocket(); // Connect on component mount

    // Cleanup function to close WebSocket on component unmount
    return () => {
      ws.current?.close();
    };
  }, [connectWebSocket]); // Depend on connectWebSocket

  // Effect to handle keyboard input for snake direction and restart
  useEffect(() => {
    const handleKeyDown = (event) => {
      if (!ws.current || ws.current.readyState !== WebSocket.OPEN) {
        // Handle reconnect/restart only if disconnected or game over
        if (event.key === 'r' || event.key === 'R') {
           if (!isConnected && !ws.current) { // Only reconnect if fully disconnected
             console.log('Reconnecting...');
             setMessage('Reconnecting...');
             connectWebSocket();
           } else if (gameState?.gameState === 'Lost') {
             // If connected but game is lost, closing and reopening triggers a new game on backend
             console.log('Restarting game...');
             setMessage('Restarting...');
             ws.current.close(); // Server will detect close and cleanup, new connection starts new game
           }
        }
        return; // Don't send directions if not connected and running
      }

      let direction = null;
      switch (event.key) {
        case 'ArrowUp':
        case 'w':
          direction = 'UP';
          break;
        case 'ArrowDown':
        case 's':
          direction = 'DOWN';
          break;
        case 'ArrowLeft':
        case 'a':
          direction = 'LEFT';
          break;
        case 'ArrowRight':
        case 'd':
          direction = 'RIGHT';
          break;
        default:
          return; // Ignore other keys
      }

      if (direction && gameState?.gameState === 'Running') {
         // Prevent default browser scroll on arrow keys
         event.preventDefault();
         // console.log(`Sending direction: ${direction}`); // Debugging
         ws.current.send(direction);
      }
    };

    window.addEventListener('keydown', handleKeyDown);

    // Cleanup function to remove event listener
    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };
  }, [isConnected, gameState, connectWebSocket]); // Re-run if connection or game state changes

  return (
    <div className="App">
      <h1>Snake Game</h1>
      <div className="game-info">
        <span>{message}</span>
        <span>High Score: {highScore}</span>
      </div>
      {gameState ? (
        <GameBoard
          boardWidth={gameState.boardWidth}
          boardHeight={gameState.boardHeight}
          snakeBody={gameState.snakeBody}
          food={gameState.food}
        />
      ) : (
        <p>Loading game board...</p> // Show loading message until first state arrives
      )}
    </div>
  );
}

export default App;
