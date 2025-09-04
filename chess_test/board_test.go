package chess_test

import (
	"testing"

	"testchess/chess"
)

func TestBoardInitialization(t *testing.T) {
	g := chess.NewGame()
	b := g.Board

	// Kings
	if p := b.At(chess.Pos{R: 0, C: 4}); p == nil || p.Type != chess.King || p.Color != chess.White {
		t.Fatalf("expected White King at e1")
	}
	if p := b.At(chess.Pos{R: 7, C: 4}); p == nil || p.Type != chess.King || p.Color != chess.Black {
		t.Fatalf("expected Black King at e8")
	}

	// Pawns
	for c := 0; c < 8; c++ {
		if p := b.At(chess.Pos{R: 1, C: c}); p == nil || p.Type != chess.Pawn || p.Color != chess.White {
			t.Fatalf("expected White Pawn at row 2 col %d", c)
		}
		if p := b.At(chess.Pos{R: 6, C: c}); p == nil || p.Type != chess.Pawn || p.Color != chess.Black {
			t.Fatalf("expected Black Pawn at row 7 col %d", c)
		}
	}
	g.Board.Print()
}
