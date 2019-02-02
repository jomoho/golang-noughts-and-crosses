package xo

import (
    "encoding/json"
    "net/http"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
)

func StartGameView(w http.ResponseWriter, r *http.Request) {
	new_game := Game{len(games), makeField(), rand.Intn(2), 0, "none"}
	games = append(games, new_game)
	json.NewEncoder(w).Encode(new_game)
}

func MoveView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	game_id, err := strconv.Atoi(params["gameId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	player_id, err := strconv.Atoi(params["playerId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	x, err := strconv.Atoi(params["x"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	y, err := strconv.Atoi(params["y"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if player_id >= 2 {
		http.Error(w, "Player with id " +  params["playerId"] + " does not exist.", 403)
		return
	}

	if game_id >= len(games) {
		http.Error(w, "Game with id " +  params["gameId"] + " does not exist.", 404)
		return
	}

	var game = games[game_id]
	if game.Winner != "none" {
		json.NewEncoder(w).Encode(game)
		return
	}

	if (game.Round + game.StartPlayer) % 2 != player_id {
		http.Error(w, "Player " +  params["playerId"] + "not your time to move", 403)
	}

	if validMove(game.Field, x, y) {
		game.Field = makeFieldMove(game.Field, tokens[player_id], x, y)
		game.Round = game.Round + 1
		game.Winner = detectWinner(game.Field)

		games[game_id] = game
		json.NewEncoder(w).Encode(game)
	} else {
        http.Error(w, "Not a valid move.", 403)
	}
}


func StateView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	game_id, err := strconv.Atoi(params["gameId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if game_id >= len(games) {
		http.Error(w, "Game with id " + params["gameId"] + " does not exist", 404)
		return
	}

	json.NewEncoder(w).Encode(games[game_id])
}
