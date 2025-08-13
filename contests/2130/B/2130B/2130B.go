package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printArray(arr []int) {
	strArr := make([]string, len(arr))

	for i, v := range arr {
		strArr[i] = strconv.Itoa(v)
	}

	result := strings.Join(strArr, " ")
	fmt.Println(result)
}

func customSort(nums []int) []int {
    // Count occurrences
    count0, count1, count2 := 0, 0, 0
    for _, num := range nums {
        switch num {
        case 0:
            count0++
        case 1:
            count1++
        case 2:
            count2++
        }
    }

    // Overwrite the array based on required order: 0s -> 2s -> 1s
    index := 0
    for i := 0; i < count0; i++ {
        nums[index] = 0
        index++
    }
    for i := 0; i < count2; i++ {
        nums[index] = 2
        index++
    }
    for i := 0; i < count1; i++ {
        nums[index] = 1
        index++
    }

    return nums
}

func sum(arr []int) int {
	ans := 0
	for _, n := range arr {
		ans += n
	}

	return ans
}

func pathless(_ int, s int, arr []int) {
	base_sum := sum(arr)

	if s < base_sum {
		printArray(arr)
	} else if base_sum == s {
		fmt.Println(-1)
	} else {
		d := s - base_sum
		if d != 1 {
			fmt.Println(-1)
		} else {
			arr = customSort(arr)
			printArray(arr)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)

	//loop over t
	
	for ; t > 0; t-- {
		// get each testcase specific data;
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		arr := strings.Split(temp, " ")
		n, _ := strconv.Atoi(arr[0])
		s, _ := strconv.Atoi(arr[1])
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		arr = strings.Split(temp, " ")
		arrconv := make([]int, len(arr))

		for i, c := range arr {
			num, _ := strconv.Atoi(c)
			arrconv[i] = num
		}

		pathless(n, s, arrconv)
	}
}