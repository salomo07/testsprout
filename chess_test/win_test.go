package chess_test

import (
	"fmt"
	"testing"

	"testchess/chess"
)

// helper: taruh bidak berdasarkan square string (misal "d1")
func placePiece(g *chess.Game, square string, pieceType chess.PieceType, color chess.Color) {
	row, col := chess.ParseSquare(square)
	g.Board.Squares[row][col] = &chess.Piece{Type: pieceType, Color: color}
}

func TestWinByKingCapture(t *testing.T) {
	g := chess.NewCustomGame()

	// Tempatkan White Queen di d1
	placePiece(g, "d1", chess.Queen, chess.Black)
	// placePiece(g, "f3", chess.King, chess.White)
	// g.Board.Squares[0][3] = &chess.Piece{Type: chess.Queen, Color: chess.White}

	// Tempatkan Black King di d8
	// g.Board.Squares[7][3] = &chess.Piece{Type: chess.King, Color: chess.Black}

	// Print board untuk verifikasi
	g.Board.Print()

	// Lakukan move d1 → e8
	mv, _ := chess.ParseMove("d1 f3")
	captured, err := g.MakeMove(mv)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if captured != nil && captured.Type == chess.King {
		fmt.Println("Black King captured! White wins.")
	}

}
