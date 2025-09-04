package chess

import (
	"errors"
	"fmt"
	"strings"
)

type Board struct {
	Squares [8][8]*Piece // [row][col]
}

func NewBoard() *Board {
	b := &Board{}
	b.setup()
	return b
}

func (b *Board) setup() {
	// Place Pawns
	for c := 0; c < 8; c++ {
		b.Squares[1][c] = &Piece{Type: Pawn, Color: White}
		b.Squares[6][c] = &Piece{Type: Pawn, Color: Black}
	}
	// Rooks
	b.Squares[0][0] = &Piece{Type: Rook, Color: White}
	b.Squares[0][7] = &Piece{Type: Rook, Color: White}
	b.Squares[7][0] = &Piece{Type: Rook, Color: Black}
	b.Squares[7][7] = &Piece{Type: Rook, Color: Black}

	// Knights
	b.Squares[0][1] = &Piece{Type: Knight, Color: White}
	b.Squares[0][6] = &Piece{Type: Knight, Color: White}
	b.Squares[7][1] = &Piece{Type: Knight, Color: Black}
	b.Squares[7][6] = &Piece{Type: Knight, Color: Black}

	// Bishops
	b.Squares[0][2] = &Piece{Type: Bishop, Color: White}
	b.Squares[0][5] = &Piece{Type: Bishop, Color: White}
	b.Squares[7][2] = &Piece{Type: Bishop, Color: Black}
	b.Squares[7][5] = &Piece{Type: Bishop, Color: Black}

	// Queens
	b.Squares[0][3] = &Piece{Type: Queen, Color: White}
	b.Squares[7][3] = &Piece{Type: Queen, Color: Black}

	// Kings
	b.Squares[0][4] = &Piece{Type: King, Color: White}
	b.Squares[7][4] = &Piece{Type: King, Color: Black}
}

func (b *Board) At(p Pos) *Piece {
	if !p.InBounds() {
		return nil
	}
	return b.Squares[p.R][p.C]
}

func (b *Board) Move(from, to Pos) (captured *Piece, err error) {
	if !from.InBounds() || !to.InBounds() {
		return nil, errors.New("out of bounds")
	}
	p := b.At(from)
	if p == nil {
		return nil, errors.New("no piece at source")
	}
	if sameColor(p, b.At(to)) {
		return nil, errors.New("cannot capture own piece")
	}
	captured = b.At(to)
	// perform move
	b.Squares[to.R][to.C] = p
	b.Squares[from.R][from.C] = nil
	return captured, nil
}

func (b *Board) Empty(p Pos) bool {
	return b.At(p) == nil
}

func (b *Board) PathClear(from, to Pos) bool {
	dr := sign(to.R - from.R)
	dc := sign(to.C - from.C)
	r := from.R + dr
	c := from.C + dc
	for r != to.R || c != to.C {
		if b.Squares[r][c] != nil {
			return false
		}
		r += dr
		c += dc
	}
	return true
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func (b *Board) Print() {
	var sb strings.Builder
	for r := 7; r >= 0; r-- {
		sb.WriteString(fmt.Sprintf("%d | ", r+1))
		for c := 0; c < 8; c++ {
			p := b.Squares[r][c]
			if p == nil {
				sb.WriteString(". ")
			} else {
				sb.WriteString(p.Symbol() + " ")
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("    a b c d e f g h\n\n")
	fmt.Print(sb.String())
}
func ParseSquare(s string) (int, int) {
	file := int(s[0] - 'a') // kolom a–h → 0–7
	rank := int(s[1] - '1') // rank 1–8 → 0–7
	row := rank             // rank 1 = row 0, rank 8 = row 7
	return row, file
}
