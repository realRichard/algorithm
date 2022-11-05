package bt

import (
	"strings"
)

// 力扣第 51 题
// N 皇后

// https://leetcode.cn/problems/n-queens/description/?languageTags=golang

// https://labuladong.github.io/algo/1/8/#二n-皇后问题

// 保存结果
var result [][][]byte

// 输入棋盘边长 n，返回所有合法的放置
func SolveNQueens(n int) [][]string {
	// 每跑一个 case 清空一下结果集，不然多个 case 结果集累积
	result = make([][][]byte, 0)
	// 表示一个棋盘
	board := make([][]byte, n)
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘
	for index := range board {
		board[index] = []byte(strings.Repeat(".", n))
	}
	solveNQueens(board, 0)
	res := make([][]string, 0)
	for _, r := range result {
		foo := make([]string, 0)
		for _, item := range r {
			foo = append(foo, string(item))
		}
		res = append(res, foo)
	}
	return res
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func solveNQueens(board [][]byte, row int) {
	// 触发结束条件
	if row == len(board) {
		b := make([][]byte, 0)
		for _, item := range board {
			bar := make([]byte, len(item))
			copy(bar, item)
			b = append(b, bar)
		}
		result = append(result, b)
		return
	}
	for col := 0; col < len(board); col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		board[row][col] = []byte("Q")[0]
		// 进入下一行决策
		solveNQueens(board, row+1)
		// 撤销选择
		board[row][col] = []byte(".")[0]
	}
}

// 是否可以在 board[row][col] 放置皇后？
func isValid(board [][]byte, row int, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for r := 0; r <= row; r++ {
		if board[r][col] == []byte("Q")[0] {
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for r, c := row, col; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == []byte("Q")[0] {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for r, c := row, col; r >= 0 && c < n; r, c = r-1, c+1 {
		if board[r][c] == []byte("Q")[0] {
			return false
		}
	}
	return true
}
