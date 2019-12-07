package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var program []int
	var full_input string

	for scanner.Scan() {
		full_input = scanner.Text()
	}

	split_input := strings.Split(full_input, ",")
	var prog_length = len(split_input)

	for i := 0; i < prog_length; i++ {
		val, err := strconv.Atoi(split_input[i])

		if err == nil {
			program = append(program, val)
		}
	}

	// mess with state as per instructions
	program[1] = 12
	program[2] = 2

	for i := 0; program[i] != 99; i += 4 {
		opcode := program[i]
		source1 := program[i+1]
		source2 := program[i+2]
		dest := program[i+3]

		switch opcode {
		case 1:
			program[dest] = program[source1] + program[source2]
		case 2:
			program[dest] = program[source1] * program[source2]
		}
	}

	fmt.Println("Final val: ", program[0])
}
