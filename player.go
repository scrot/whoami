package main

// interacter provides an interface describing the
// game interaction required, it can be implemented
// to provide a custom interaction
type interacter interface {
	playerName() string
	toGuess(active player, target player) string
	guess(active player) string
}

type player struct {
	index    int
	name     string
	target   string
	score    int
	interact interacter
}
