# Noughts and Crosses (Tic-Tac-Toe) in Go

## Overview

This application is a simple implementation of the Noughts and Crosses (Tic-Tac-Toe) game using Go. It utilizes goroutines, channels, and sync wait groups to manage concurrent operations, allowing for a smooth and interactive gameplay experience.

## Techniques used

- Concurrency
- Go Routines
- Channels
- sync waitgroups

  
## Features

- Two players can take turns to play the game.
- The game board is displayed in the console.
- Players input their moves through the console.
- The game detects and announces the winner or if the game is a draw.

## Requirements

- Go 1.18 or later

## Files

- `main.go`: Contains the implementation of the game logic, game loop, user input handling, and board rendering.

## How to Run

1. **Clone the Repository**

   ```sh
   git clone <repository-url>
   cd <repository-directory>
## Build the Application

go build -o tic-tac-toe

## Run the Application

go run main.go

