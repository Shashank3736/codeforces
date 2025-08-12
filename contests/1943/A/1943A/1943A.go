package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mexGame(hash []int, n int) int {
	turn := true
	// c := make(map[int]bool)
	for i := 0; i < n; {
		// fmt.Println(i, turn, hash[i])
		if turn {
			// alice chance
			if hash[i] == 0 {
				return i
			}
			if hash[i] > 1 {
				i++
				continue
			}
			hash[i]--
			i++
			turn = false
		} else {
			// bob chance
			if hash[i] > 1 {
				i++
				continue
			} else {
				return i
			}
		}
	}
	return n
}

// func copyCatMex(nums []int, n int) int {
// 	c := 0
// 	for i := 0; i < n; i++ {
// 		if nums[i] == 1 {
// 			c++
// 		}
// 		if c == 2 || nums[i] == 0 {
// 			return i
// 		}
// 	}
// 	return 0
// }

func solve(scanner *bufio.Scanner) {
	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)
	scanner.Scan()
	s := scanner.Text()
	// s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	nums := make([]int, n+1)

	// if n > 7250 {
	// 	fmt.Println(n, s)
	// }
	
	for _, v := range arr {
		num, _ := strconv.Atoi(v)
		// if num > n {
		// 	continue
		// }
		nums[num]++
	}
	fmt.Println(mexGame(nums, n))
}

func main() {
	// check something here
	r := os.Stdin
	if f, err := os.Open("test.txt"); err == nil {
		r = f
		defer f.Close()
	}
	// scan cmd
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, bufio.MaxScanTokenSize), 10*2*100*1000)
	scanner.Scan()
	var t int
	fmt.Sscanf(scanner.Text(), "%d", &t)
	// get input from each testcase
	for ; t > 0; t-- {
		solve(scanner)
	}
}
