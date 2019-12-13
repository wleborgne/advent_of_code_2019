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

	// count orbits from initial point (YOU, SAN) to COM
	common_body := find_common_body("YOU", "SAN", orbit_list)
	transfers := count_orbits("YOU", common_body, orbit_list) + count_orbits("SAN", common_body, orbit_list) - 2
	// difference between counts for origins - 2 ? NO

	fmt.Println("total transfers: ", transfers)
}

func count_orbits(start_body string, dest_body string, orbits map[string]string) int {
	i := 0
	next_key := start_body

	for next_key != dest_body {
		next_key = orbits[next_key]
		i++
	}

	return i
}

func find_common_body(first_body string, second_body string, orbits map[string]string) string {
	first_list := make_orbit_list(first_body, orbits)
	second_list := make_orbit_list(second_body, orbits)

	for i := range first_list {
		for j := range second_list {
			if first_list[i] == second_list[j] {
				return first_list[i]
			}
		}
	}
	return "FOO"
}

func make_orbit_list(body string, orbits map[string]string) []string {
	var orbit_list []string
	var next_key = body

	for next_key != "COM" {
		orbit_list = append(orbit_list, orbits[next_key])
		next_key = orbits[next_key]
	}

	return orbit_list
}
