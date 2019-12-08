package main

import (
	"fmt"
	"strconv"
)

func main() {
	var results []int

	for i := 138307; i <= 654504; i++ {
		s := strconv.Itoa(i)
		if increase(s) && doubles(s) {
			results = append(results, i)
		}
	}

	fmt.Println("Number of possible values: ", len(results))
}

func increase(s string) bool {
	return s[0] <= s[1] && s[1] <= s[2] && s[2] <= s[3] && s[3] <= s[4] && s[4] <= s[5]
}

func doubles(s string) bool {
	return s[0] == s[1] || s[1] == s[2] || s[2] == s[3] || s[3] == s[4] || s[4] == s[5]
}
