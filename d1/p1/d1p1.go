package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(calculate_fuel(os.Stdin))
}

func calculate_fuel(r io.Reader) int {
	var total = 0
	var scanner = bufio.NewScanner(r)

	for scanner.Scan() {
		total += calculate_unit(scanner.Text())
	}

	return total
}

func calculate_unit(s string) int {
	val, err := strconv.Atoi(s)

	if err == nil {
		val = (int(val / 3.0)) - 2
		return val
	} else {
		return 0
	}
}
