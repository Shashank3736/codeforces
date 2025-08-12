package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func leftAndRightHouses(n int, a []string) int {
	mid := float64(n)/2.0
	total := 0

	for i:=0; i<n; i++ {
		if a[i] == "1" {
			total++
		} else {
			total--
		}
	}

	cur := 0
	var ans int
	if total >= 0 {
		ans = 0
	} else {
		ans = n
	}

	for i := 0; i < n; i++ {
		if a[i] == "1" {
			cur++
			total--
		} else {
			cur--
			total++
		}

		if cur <= 0 && total >= 0 {
			if abs(mid-float64(i+1)) < abs(mid-float64(ans)) {
				ans = i+1
			}
		}
	}

	return ans
}

func main() {
	// get testcases t
	reader := bufio.NewReader(os.Stdin)
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	t, _ := strconv.Atoi(temp)
	// loop over it
	for ; t > 0; t-- {
		// get int n where n > 2
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		n, _ := strconv.Atoi(temp)
		// get string of length a
		temp, _ = reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		a := strings.Split(temp, "")
		ans := leftAndRightHouses(n, a)
		fmt.Println(ans)
	}
}