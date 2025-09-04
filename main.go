package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"testchess/chess"
)

func main() {
	game := chess.NewGame()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Console Chess (Go) ===")
	fmt.Println("Input move format:")
	// fmt.Println("- Numeric: 1,3 2,3  (row,col row,col)")
	fmt.Println("- Algebraic: b2 b3   (fileRank fileRank)")
	fmt.Println("White starts. Game ends when a King is captured.\n")

	for {
		game.Board.Print()

		active := "White"
		if game.Turn == chess.Black {
			active = "Black"
		}
		fmt.Printf("%s to move. Enter move: ", active)

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.EqualFold(line, "exit") || strings.EqualFold(line, "quit") {
			fmt.Println("Goodbye!")
			return
		}

		mv, err := chess.ParseMove(line)
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}

		captured, err := game.MakeMove(mv)
		if err != nil {
			fmt.Println("Move error:", err)
			continue
		}

		if captured != nil && captured.Type == chess.King {
			game.Board.Print()
			fmt.Printf("%s captured the King. %s wins!\n", active, active)
			return
		}

		game.ToggleTurn()
	}
}
