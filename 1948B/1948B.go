package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arrayFix(nums []int) bool {
	prev := 0

	for _, num := range nums {
		if prev > num {
			return false
		}

		tenth := num / 10
		ones := num % 10

		if prev <= tenth && tenth <= ones {
			prev = ones
		} else {
			prev = num
		}
	}
	return true
}
func main() {
	// get no. of testcases t
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)
	// loop over it
	for ; t > 0; t-- {
		// get size of array
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		n, _ := strconv.Atoi(temp)
		// get array elements
		arr := make([]int, n)
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		tempArr := strings.Split(temp, " ")
		for i := 0; i < n; i++ {
			arr[i], _ = strconv.Atoi(tempArr[i])
		}
		// check if array can be fixed
		if arrayFix(arr) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}