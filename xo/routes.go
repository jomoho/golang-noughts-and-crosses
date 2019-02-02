package xo

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/start_game", StartGameView).Methods("POST")
	router.HandleFunc("/move", MoveView).Queries(
			"gameId", "{gameId}",
			"playerId", "{playerId}",
			"x", "{x:[0-2]+}",
			"y", "{y:[0-2]+}").Methods("POST")
	router.HandleFunc("/state", StateView).Queries(
			"gameId", "{gameId}").Methods("GET")

	return router
}