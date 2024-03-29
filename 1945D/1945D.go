package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func seraphimTheOwl(coins [][]int, m int) int {
	bTotal := 0
	res := 0
	for i := len(coins)-1; i >= m; i-- {
		bTotal += coins[i][1]
		if coins[i][1] > coins[i][0] {
			res += bTotal + coins[i][0] - coins[i][1]
			bTotal = 0
		}
	}

	temp := bTotal + coins[m-1][0]

	for i := m-1; i >= 0; i-- {
		if coins[i][0] + bTotal < temp {
			temp = coins[i][0] + bTotal
		}
		bTotal += coins[i][1]
	}

	return res + temp
	
}

func main() {
	// get no. of testcases t
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)
	// loop over it
	for ; t > 0; t-- {
		// get int n and m as n m input
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		tempArr := strings.Split(temp, " ")
		n, _ := strconv.Atoi(tempArr[0])
		m, _ := strconv.Atoi(tempArr[1])
		// get array of int
		arr := make([][]int, n)
		// get array a
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		tempArrA := strings.Split(temp, " ")
		// get array b
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		tempArrB := strings.Split(temp, " ")
		// loop over n
		for i := 0; i < n; i++ {
			arr[i] = make([]int, 2)
			arr[i][0], _ = strconv.Atoi(tempArrA[i])
			arr[i][1], _ = strconv.Atoi(tempArrB[i])
		}
		// get answer
		fmt.Println(seraphimTheOwl(arr, m))
	}
}
