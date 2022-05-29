// https://leetcode-cn.com/problems/valid-tic-tac-toe-state/
package main

import "fmt"

/**
 * 找出的三条规则
 * 1. X 数量num(X) 要么等 num(O), 要么num(X) = num(O) + 1
 * 2. 当 X 有直线相同字符时, num(X) = num(O) + 1; 当 O 有直线相同字符时, num(X) = num(O);
 * 3. X O 不会同时有直线相同字符
 */
func validTicTacToe(board []string) bool {
	num := map[byte]int{
		'X': 0,
		'O': 0,
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == ' ' {
				continue
			}
			num[board[i][j]]++
		}
	}
	// 规则1
	if num['X'] != num['O'] && num['X'] != num['O']+1 {
		return false
	}

	winX, winO := 0, 0
	for i := 0; i < 3; i++ {
		if board[i][0] == 'X' && board[i][0] == board[i][1] && board[i][0] == board[i][2] ||
			board[0][i] == 'X' && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			winX++
		}
		if board[i][0] == 'O' && board[i][0] == board[i][1] && board[i][0] == board[i][2] ||
			board[0][i] == 'O' && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			winO++
		}
	}
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		if board[0][0] == 'X' {
			winX++
		} else if board[0][0] == 'O' {
			winO++
		}
	}
	if board[2][0] == board[1][1] && board[2][0] == board[0][2] {
		if board[2][0] == 'X' {
			winX++
		} else if board[2][0] == 'O' {
			winO++
		}
	}
	// 规则2、3
	if winX >= 1 && (num['X'] != num['O']+1 || winO >= 1) || winO >= 1 && num['X'] != num['O'] {
		return false
	}
	return true
}

func main() {
	board := [][]string{
		{"O  ", "   ", "   "},
		{"XOX", " X ", "   "},
		{"XXX", "   ", "OOO"},
		{"XOX", "O O", "XOX"},
		{"XXX", "OOX", "OOX"},
		{"XXX", "XOO", "OO "},
	}

	for _, str := range board {
		fmt.Println("ret:", validTicTacToe(str))
	}
}
