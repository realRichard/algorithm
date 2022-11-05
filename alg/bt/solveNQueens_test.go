package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	result := bt.SolveNQueens(4)
	fmt.Println(result)
}
