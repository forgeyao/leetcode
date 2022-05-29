// https://leetcode-cn.com/problems/determine-color-of-a-chessboard-square/
package main

import "fmt"

func squareIsWhite(coordinates string) bool {
	return (coordinates[0]%2 != coordinates[1]%2)
}

func main() {
	s := []string{"a1", "h3", "c7"}
	for _, ss := range s {
		fmt.Println("ret:", squareIsWhite(ss))
	}
}
