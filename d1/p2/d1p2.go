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
		total += calculate_cargo(scanner.Text())
	}

	return total
}

func calculate_cargo(s string) int {
	val, err := strconv.Atoi(s)

	if err == nil {
		var mod_fuel = get_fuel(val)
		return calculate_total(mod_fuel)
	} else {
		return 0
	}
}

func get_fuel(v int) int {
	return (int(v/3.0) - 2)
}

func calculate_total(cargo_fuel int) int {
	var total = 0

	for inc_fuel := cargo_fuel; inc_fuel > 0; inc_fuel = get_fuel(inc_fuel) {
		total += inc_fuel
	}

	return total
}
