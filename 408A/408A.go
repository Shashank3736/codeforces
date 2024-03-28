package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := 0
	// get no. of cashier t
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)
	// loop it over
	arr := make([]int, t)
	// fill array with no. of people at each cashier
	temp, _ = reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	tempArr := strings.Split(temp, " ")

	for i := 0; i < t; i++ {
		arr[i], _ = strconv.Atoi(tempArr[i])
	}

	// calculate timings

	for _, ps := range arr {
		temp_ans := ps*15

		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		tempArr = strings.Split(temp, " ")

		for i := 0; i < ps; i++ {
			items, _ := strconv.Atoi(tempArr[i])
			temp_ans += items*5
		}

		if ans == 0 || temp_ans < ans {
			ans = temp_ans
		}
	}

	fmt.Println(ans)
}
