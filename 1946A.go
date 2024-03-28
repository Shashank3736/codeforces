package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Change median of an array by using minimum amount of methods.
func medianOfAnArray(arr []int) int {
	// sort array
	sort.Ints(arr)
	// get length of array
	n := len(arr)
	// get position of median
	i := n / 2

	if n % 2 == 0 {
		i--
	}

	value := arr[i]
	ans := 1

	i++
	for ; i < n; i++ {
		if arr[i] != value {
			break
		}
		ans++
	}

	return ans
}

func main() {
	// get no. of testcases t
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)

	// loop it over
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
		// get no. of methods to change median
		ans := medianOfAnArray(arr)
		fmt.Println(ans)
	}
}
