package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maximumSumOfSubArray(arr []int) int {
	ans := 0

	for i := 0; i < len(arr); i++ {
		temp := 0
		if arr[i] < 1 {
			continue
		}
		for j := i; j < len(arr); j++ {
			temp += arr[j]
			if temp > ans {
				ans = temp
			}
			if temp <= 0 {
				break
			}
		}
	}

	return ans
}

func maximumSum(nums []int, k int) int {
	ans := 0

	for _, num := range nums {
		ans += num
	}

	subArrayTotal := maximumSumOfSubArray(nums)
	
	ans += subArrayTotal
	for i := 1; i < k; i++ {
		subArrayTotal *= 2
		subArrayTotal %= 1000000007
		ans += subArrayTotal
		ans %= 1000000007
		
	}
	
	ans %= 1000000007

	if ans < 0 {
		ans += 1000000007
	}

	return ans
}

func main() {
	// get no. of testcases t
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)
	// loop over it
	for ; t > 0; t-- {
		// get int n and k as n k input
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		tempArr := strings.Split(temp, " ")
		n, _ := strconv.Atoi(tempArr[0])
		k, _ := strconv.Atoi(tempArr[1])

		// get array elements
		nums := make([]int, n)
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		tempArr = strings.Split(temp, " ")
		for i := 0; i < n; i++ {
			nums[i], _ = strconv.Atoi(tempArr[i])
		}
		// get answer
		ans := maximumSum(nums, k)
		fmt.Println(ans)
	}
}
