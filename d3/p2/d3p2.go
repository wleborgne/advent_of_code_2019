package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var lines [2]string

	scanner.Scan()
	lines[0] = scanner.Text()
	scanner.Scan()
	lines[1] = scanner.Text()

	red_wire_path := strings.Split(lines[0], ",")
	blue_wire_path := strings.Split(lines[1], ",")

	red_points := gather_points(red_wire_path)
	blue_points := gather_points(blue_wire_path)

	common_points := get_common_points(red_points, blue_points)

	distances := get_distances(common_points, red_points, blue_points)
	sort.Ints(distances)

	fmt.Println("Shortest intersection distance: ", distances[0])
}

func gather_points(path []string) []point {
	var points []point
	var last_point point
	last_point.x = 0
	last_point.y = 0

	for i := 0; i < len(path); i++ {
		direction := string(path[i][0])
		distance, _ := strconv.Atoi(path[i][1:])

		points = append(points, points_for_segment(last_point, direction, distance)...)
		last_point = points[len(points)-1]
	}

	return points
}

func points_for_segment(start point, direction string, distance int) []point {
	var inc point
	var last = start
	var segment = make([]point, distance)

	switch direction {
	case "U":
		inc.x = 0
		inc.y = 1
	case "D":
		inc.x = 0
		inc.y = -1
	case "R":
		inc.x = 1
		inc.y = 0
	case "L":
		inc.x = -1
		inc.y = 0
	}

	for i := 1; i <= distance; i++ {
		last.x += inc.x
		last.y += inc.y
		segment[i-1] = last
	}

	return segment
}

func get_common_points(red []point, blue []point) []point {
	var common []point

	for i := 0; i < len(red); i++ {
		for j := 0; j < len(blue); j++ {
			if red[i].x == blue[j].x && red[i].y == blue[j].y {
				common = append(common, red[i])
			}
		}
	}

	return common
}

func get_distances(common []point, red []point, blue []point) []int {
	var d = make([]int, len(common))

	// for each common point
	for i := range common {
		p := common[i]
		red_index := find_point(p, red)
		blue_index := find_point(p, blue)
		d[i] = red_index + blue_index
	}

	return d
}

func find_point(p point, path []point) int {
	var index = 0

	for i := range path {
		current := path[i]
		if current.x == p.x && current.y == p.y {
			index = i
			break
		}
	}

	return index + 1
}
