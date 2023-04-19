package main

func main() {
	game := newGame(cliEngine{})
	for i := 0; i < 2; i++ {
		game = game.addPlayer()
	}

	game.start()
}
