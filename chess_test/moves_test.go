package chess_test

import (
	"testing"

	"testchess/chess"
)

func TestValidAndInvalidMoves(t *testing.T) {
	g := chess.NewGame()

	// White: move pawn e2 to e4 (two steps)
	mv, err := chess.ParseMove("e2 e4")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if _, err := g.MakeMove(mv); err != nil {
		t.Fatalf("expected legal pawn double move, got: %v", err)
	}
	g.ToggleTurn()

	// Black: invalid move knight from b8 to b6 (knight doesn't move straight)
	mv, _ = chess.ParseMove("b8 b6")
	if _, err := g.MakeMove(mv); err == nil {
		t.Fatalf("expected illegal knight move")
	}
	// Try a legal knight move: b8 -> c6
	mv, _ = chess.ParseMove("b8 c6")
	if _, err := g.MakeMove(mv); err != nil {
		t.Fatalf("expected legal knight move, got: %v", err)
	}
	g.ToggleTurn()

	// White: try bishop from f1 to b5 (path must be clear: currently pawn on e2 moved, path clear)
	mv, _ = chess.ParseMove("f1 b5")
	if _, err := g.MakeMove(mv); err != nil {
		t.Fatalf("expected legal bishop move, got: %v", err)
	}
	g.Board.Print()
}

func TestBlockPath(t *testing.T) {
	g := chess.NewGame()
	// Try rook from a1 to a3 at start (blocked by own pawn at a2)
	mv, _ := chess.ParseMove("a1 a3")
	if _, err := g.MakeMove(mv); err == nil {
		t.Fatalf("expected blocked rook move to be illegal")
	}
}
