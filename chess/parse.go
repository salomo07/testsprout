package chess

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Examples:
// "1,3 2,3" -> from (0-based: 0,2) ? No. Spec expects 1..8 indexes.
// We'll accept row,col in 1..8 then convert to 0-based.
// "b2 b3" -> ('b'=1, '2'=1) -> (1,1) etc.
func ParseMove(s string) (Move, error) {
	s = strings.TrimSpace(s)
	parts := strings.Fields(s)
	if len(parts) != 2 {
		return Move{}, errors.New("expected two coordinates separated by space")
	}

	// Try numeric "r,c"
	from, errA := parseNumeric(parts[0])
	to, errB := parseNumeric(parts[1])
	if errA == nil && errB == nil {
		return Move{From: from, To: to}, nil
	}

	// Try algebraic "b2"
	from, errA = parseAlgebraic(parts[0])
	to, errB = parseAlgebraic(parts[1])
	if errA == nil && errB == nil {
		return Move{From: from, To: to}, nil
	}

	return Move{}, fmt.Errorf("invalid input; numeric error: %v, algebraic error: %v", errA, errB)
}

var reNum = regexp.MustCompile(`^\s*(\d)\s*,\s*(\d)\s*$`)

func parseNumeric(token string) (Pos, error) {
	m := reNum.FindStringSubmatch(token)
	if m == nil {
		return Pos{}, errors.New("numeric format must be r,c (digits 1-8)")
	}
	r, _ := strconv.Atoi(m[1])
	c, _ := strconv.Atoi(m[2])
	if r < 1 || r > 8 || c < 1 || c > 8 {
		return Pos{}, errors.New("row/col must be 1..8")
	}
	// Convert to 0-based internal (rows: 1 bottom -> index 0 ; but we print inverse)
	// We'll interpret row "1" as rank 1 (bottom, White side) -> internal row index 0.
	return Pos{R: r - 1, C: c - 1}, nil
}

func parseAlgebraic(token string) (Pos, error) {
	token = strings.ToLower(strings.TrimSpace(token))
	if len(token) != 2 {
		return Pos{}, errors.New("algebraic must be like 'b2'")
	}
	file := token[0] // 'a'..'h'
	rank := token[1] // '1'..'8'
	if file < 'a' || file > 'h' || rank < '1' || rank > '8' {
		return Pos{}, errors.New("file must be a..h and rank 1..8")
	}
	c := int(file - 'a')
	r := int(rank - '1') // '1' -> 0
	return Pos{R: r, C: c}, nil
}
