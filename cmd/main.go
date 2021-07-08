package main

import (
	"bufio"
	"fmt"
	"github.com/dihmuzikien/bowling"
	"os"
	"strings"
)

type gameStateViewer interface {
	Score() []int
	CurrentFrame() int
}

func main() {
	g := bowling.NewGame()
	for !g.Finished() {
		currentFrame := g.CurrentFrame()
		play, err := input(currentFrame)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("scoring: <%s> for frame %d\n", play, currentFrame)
		playErr := g.Play(play)
		if playErr != nil {
			fmt.Printf("failed to score: %v\n", playErr)
			continue
		}
		fmt.Printf("finished frame #%d\n", currentFrame)
		printGameState(g)
	}
}

func printGameState(g gameStateViewer) {
	fmt.Printf("Current Score is %v\n", g.Score())
	fmt.Println("----------------------------")
}

func input(frame int) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("record your score for frame %d: ", frame)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	strTrim := strings.Trim(userInput, "\t \n")
	return strTrim, nil
}
