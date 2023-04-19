package main

import (
	"errors"
	"math/rand"
	"time"
)

var (
	errNoActivePlayer = errors.New("active player not set")
)

type game struct {
	active  *player
	players []player
}

// newGame creates an empty game and init rand.
func newGame() game {
	rand.Seed(time.Now().UnixNano())
	return game{nil, []player{}}
}

// listen for players instead of adding them
func (g game) addPlayer() game {
	p := player{
		index: len(g.players),
		name:  g.playerName(),
	}
	g.players = append(g.players, p)
	return g
}

func (g game) start() {
	//set active player
	active := g.randPlayer()

	// set guess targets for each player
	for i := range g.players {
		var (
			current = g.players[i]
			next    = g.nextPlayer(current)
		)

		next.target = g.toGuess(current, next)
	}

	g.guess(active)
}

func (g game) nextPlayer(current player) player {
	ni := current.index + 1
	if ni >= len(g.players) {
		ni = 0
	}
	return g.players[ni]
}

func (g game) randPlayer() player {
	ri := rand.Intn(len(g.players)) - 1
	return g.players[ri]
}
