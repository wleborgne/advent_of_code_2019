package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var orbit_list = make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		bodies := strings.Split(line, ")")
		orbit_list[bodies[1]] = bodies[0]
	}
	fmt.Printf("Got %d orbits in input.\n", len(orbit_list))

	var orbit_total int = 0
	for body, _ := range orbit_list {
		orbit_total += count_orbits(body, orbit_list)
	}

	fmt.Println("total orbits: ", orbit_total)
}

func count_orbits(body string, orbits map[string]string) int {
	i := 0
	next_key := body

	for next_key != "COM" {
		next_key = orbits[next_key]
		i++
	}

	return i
}
