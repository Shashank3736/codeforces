package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// arrowPath checks if there is a path from the top-left cell to the bottom-right cell of a 2D grid.
//
// n is the number of rows and columns in the grid.
// row1 is a string array representing the first row of the grid.
// row2 is a string array representing the second row of the grid.
// Returns a boolean indicating if there is a path from the top-left cell to the bottom-right cell.
func arrowPath(n int, rows [][]string) bool {
	dp := make(map[string]bool)

	var dfs func(int, int) bool

	dfs = func(i, j int) bool {
		ans := false // initial value
		// check for all conditions
		if i < 0 || i >= n || j < 0 || j >= n {
			return false
		}
		if i == 1 && j == n-1 {
			return true
		}
		if dp[fmt.Sprintf("%d %d", i, j)] {
			return false
		}
		// ==========================================================================
		// set dp value true
		dp[fmt.Sprintf("%d %d", i, j)] = true
		// top
		if i == 1 {
			if rows[i-1][j] == ">" {
				ans = ans || dfs(i-1, j+1)
			} else {
				ans = ans || dfs(i-1, j-1)
			}
		}
		// bottom
		if i == 0 {
			if rows[i+1][j] == ">" {
				ans = ans || dfs(i+1, j+1)
			} else {
				ans = ans || dfs(i+1, j-1)
			}
		}
		// left
		if j > 0 {
			if rows[i][j-1] == "<" {
				ans = ans || dfs(i, j-2)
			}
		}
		// right
		if j < n-1 {
			if rows[i][j+1] == ">" {
				ans = ans || dfs(i, j+2)
			}
		}
		return ans
	}

	return dfs(0, 0)
}

func main() {
	// get no. of testcases t
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)
	// loop it over
	for ; t > 0; t-- {
		// get value of n
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		n, _ := strconv.Atoi(temp)
		// get string row1
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		row1 := strings.Split(temp, "")
		// get string row2
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		row2 := strings.Split(temp, "")

		if arrowPath(n, [][]string{row1, row2}) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
