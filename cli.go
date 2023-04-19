package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliEngine struct{}

func (ce cliEngine) newPlayer() player {
	name := ask("What is your name?\n")
	return player{name: name}
}

func (ce cliEngine) toGuess(current player, target player) string {
	//TODO: pass person count
	toGuess := ask(fmt.Sprintf(
		"%s who should %s guess?\n",
		current.name,
		target.name,
	))
	return toGuess
}

func (ce cliEngine) guess(current player) string {
	guess := ask(fmt.Sprintf("%s who are you?\n", current.name))
	return guess
}

func ask(question string) string {
	fmt.Printf("%s", question)
	input := bufio.NewReader(os.Stdin)

	answer, err := input.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return answer[:len(answer)-1]
}
