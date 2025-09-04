package chess

import (
	"errors"
	"fmt"
)

type Game struct {
	Board *Board
	Turn  Color
}

func NewGame() *Game {
	return &Game{
		Board: NewBoard(),
		Turn:  White,
	}
}

type Move struct {
	From Pos
	To   Pos
}

func (g *Game) ToggleTurn() {
	if g.Turn == White {
		g.Turn = Black
	} else {
		g.Turn = White
	}
}

func (g *Game) MakeMove(m Move) (*Piece, error) {
	if !m.From.InBounds() || !m.To.InBounds() {
		return nil, errors.New("coordinates out of bounds")
	}
	p := g.Board.At(m.From)
	if p == nil {
		return nil, errors.New("no piece at source")
	}
	if p.Color != g.Turn {
		return nil, fmt.Errorf("it's %s's turn", g.Turn)
	}
	if sameColor(p, g.Board.At(m.To)) {
		return nil, errors.New("cannot capture own piece")
	}

	if !g.isLegalMove(p, m.From, m.To) {
		return nil, errors.New("illegal move for piece")
	}

	// Execute
	return g.Board.Move(m.From, m.To)
}

// Basic chess rules (no check/checkmate, no castling, no en passant, no promotion)
func (g *Game) isLegalMove(p *Piece, from, to Pos) bool {
	switch p.Type {
	case Pawn:
		return g.legalPawn(p, from, to)
	case Rook:
		return g.legalRook(from, to)
	case Knight:
		return g.legalKnight(from, to)
	case Bishop:
		return g.legalBishop(from, to)
	case Queen:
		return g.legalQueen(from, to)
	case King:
		return g.legalKing(from, to)
	default:
		return false
	}
}

func (g *Game) legalPawn(p *Piece, from, to Pos) bool {
	dir := 1 // White moves up rows (to higher index)
	startRow := 1
	if p.Color == Black {
		dir = -1 // Black moves down rows
		startRow = 6
	}

	dr := to.R - from.R
	dc := to.C - from.C
	target := g.Board.At(to)

	// Forward move
	if dc == 0 {
		// one step
		if dr == dir && target == nil {
			return true
		}
		// two steps from start
		if from.R == startRow && dr == 2*dir && target == nil {
			// square in between must be empty
			mid := Pos{R: from.R + dir, C: from.C}
			return g.Board.Empty(mid)
		}
		return false
	}

	// Capture diagonally
	if abs(dc) == 1 && dr == dir && target != nil && target.Color != p.Color {
		return true
	}

	return false
}

func (g *Game) legalRook(from, to Pos) bool {
	if from.R != to.R && from.C != to.C {
		return false
	}
	return g.Board.PathClear(from, to)
}

func (g *Game) legalBishop(from, to Pos) bool {
	if abs(to.R-from.R) != abs(to.C-from.C) {
		return false
	}
	return g.Board.PathClear(from, to)
}

func (g *Game) legalQueen(from, to Pos) bool {
	if from.R == to.R || from.C == to.C {
		return g.Board.PathClear(from, to)
	}
	if abs(to.R-from.R) == abs(to.C-from.C) {
		return g.Board.PathClear(from, to)
	}
	return false
}

func (g *Game) legalKnight(from, to Pos) bool {
	dr := abs(to.R - from.R)
	dc := abs(to.C - from.C)
	return (dr == 2 && dc == 1) || (dr == 1 && dc == 2)
}

func (g *Game) legalKing(from, to Pos) bool {
	dr := abs(to.R - from.R)
	dc := abs(to.C - from.C)
	return dr <= 1 && dc <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func NewEmptyGame() *Game {
	return &Game{
		Board: &Board{},
		Turn:  White,
	}
}

func NewCustomGame() *Game {
	g := &Game{Board: &Board{}, Turn: White}
	return g
}
