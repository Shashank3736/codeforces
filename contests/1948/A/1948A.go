package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func specialCharacters(n int) string {
	if n % 2 == 1 {
		return "NO"
	}
	ans := ""

	for i := 0; i < n / 2; i++ {
		if i % 2 == 0 {
			ans += "AA"
		} else {
			ans += "BB"
		}
	}

	return fmt.Sprintf("YES\n%s", ans)
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
		ans := specialCharacters(n)
		fmt.Println(ans)
	}
}