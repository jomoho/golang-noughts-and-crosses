# Golang Noughts And Crosses
This is a very simple API with routes,  for creating a new game, for making moves within a game and viewing the state of the game.

## Install
`go get github.com/gorilla/mux`

## Running
`go run main.go`

## Routes:

### StartGame
`POST http://localhost:8000/start_game`

Creates a new game and returns the initial state.

### Move
`POST http://localhost:8000/move`

Makes a move using the following query parameters and returns new state or error message.

#### Parameters:
```
gameId          // id of the created game
playerId        // 0 or 1
x               // between 0 and 3
y               // between 0 and 3

```

Player `0 is "X"` and  player `1 is "O"`.

### State
`GET http://localhost:8000/state`

Returns the current state of the game.

#### Parameters:
```
gameId          // id of the created game
```

## Game State
All routes return a valid game state
```
{
    "id": 1,                         // the gameId
    "field": "***\n***\n***\n",      // the playing field
    "start_player": 1,               // the playerId of starting player 0 or 1
    "round": 0,                      // round counter
    "winner": "none"                 // Winner "none", "X" or "O"
}

```

## Project structure
```
xo
-model.go       model definitions
-routes.go      routes definitions
-service.go     functionality for xo service
-views.go       views StartGameView & MoveView

main.go         entrypoint for the app
```
### Thoughts about this structure
I put `xo` into a separate package to keep things modular.
I am splitting the code, losely following a MVC pattern with an added separation of the routes definitions from the view definitions.

`service.go` contains the application state and the functions that manipulate state. Which in a real world scenario could be further seperated into a database connection and queries for example.

There is some direct manipulation of state in `views.go`, which could be further refactored to reside only in `service.go`.

### Optimization options
There is room for optimization as the state gets copied a few times this could be avoided, e.g. by using pointers.