package chess

import "fmt"

type Color int

const (
	White Color = iota
	Black
)

func (c Color) String() string {
	if c == White {
		return "White"
	}
	return "Black"
}

type PieceType int

const (
	Pawn PieceType = iota
	Rook
	Knight
	Bishop
	Queen
	King
)

type Piece struct {
	Type  PieceType
	Color Color
}

func (p *Piece) Symbol() string {
	var s string
	switch p.Type {
	case Pawn:
		s = "P"
	case Rook:
		s = "R"
	case Knight:
		s = "N"
	case Bishop:
		s = "B"
	case Queen:
		s = "Q"
	case King:
		s = "K"
	default:
		s = "?"
	}
	if p.Color == Black {
		return s // lowercase optional; keep uppercase for clarity
	}
	return s
}

func sameColor(a, b *Piece) bool {
	return a != nil && b != nil && a.Color == b.Color
}

func oppColor(a, b *Piece) bool {
	return a != nil && b != nil && a.Color != b.Color
}

type Pos struct {
	R int // 0..7 (rank index)
	C int // 0..7 (file index)
}

func (p Pos) InBounds() bool {
	return p.R >= 0 && p.R < 8 && p.C >= 0 && p.C < 8
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.R, p.C)
}
