import React from 'react';
import './GameBoard.css';

const CELL_SIZE = 20; // Size of each cell in pixels

const GameBoard = ({ boardWidth, boardHeight, snakeBody, food }) => {
  const cells = [];

  // Create grid cells (optional, for background styling)
  // for (let y = 0; y < boardHeight; y++) {
  //   for (let x = 0; x < boardWidth; x++) {
  //     cells.push(<div key={`${x}-${y}`} className="cell" style={{ left: `${x * CELL_SIZE}px`, top: `${y * CELL_SIZE}px` }}></div>);
  //   }
  // }

  // Render snake segments
  const snakeElements = snakeBody.map((segment, index) => (
    <div
      key={`snake-${index}`}
      className="snake-segment"
      style={{
        left: `${segment.X * CELL_SIZE}px`,
        top: `${segment.Y * CELL_SIZE}px`,
        width: `${CELL_SIZE}px`,
        height: `${CELL_SIZE}px`,
      }}
    ></div>
  ));

  // Render food
  const foodElement = food ? (
    <div
      className="food"
      style={{
        left: `${food.X * CELL_SIZE}px`,
        top: `${food.Y * CELL_SIZE}px`,
        width: `${CELL_SIZE}px`,
        height: `${CELL_SIZE}px`,
      }}
    ></div>
  ) : null;

  return (
    <div
      className="game-board"
      style={{
        width: `${boardWidth * CELL_SIZE}px`,
        height: `${boardHeight * CELL_SIZE}px`,
        position: 'relative', // Needed for absolute positioning of snake/food
      }}
    >
      {/* {cells} */}
      {snakeElements}
      {foodElement}
    </div>
  );
};

export default GameBoard;