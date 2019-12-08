package main

import (
	"fmt"
	"strconv"
)

func main() {
	var results []int

	for i := 138307; i <= 654504; i++ {
		s := strconv.Itoa(i)
		if increase(s) && doubles(s) && strict(s) {
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

func strict(s string) bool {
	// there must be one digit _exactly_ twice somewhere?
	var count = map[byte]int{'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0}

	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	for _, val := range count {
		if val == 2 {
			return true
		}
	}

	return false
}
