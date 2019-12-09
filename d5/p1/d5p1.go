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

	run_program(program, 1)
}

func run_program(program []int, input int) int {
	var opcode, source1, source2, dest int
	max_index := len(program) - 1
	step := 1

	for i := 0; i <= max_index-3 && program[i] != 99; i += step {

		opcode = program[i]
		//		fmt.Println("Instruction at position ", i, ": ", opcode)
		switch opcode {
		// addition: 1, 101, 1001, 1101
		case 1:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = program[source1] + program[source2]
		case 101:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = source1 + program[source2]
		case 1001:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = program[source1] + source2
		case 1101:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = source1 + source2
		// multiplication: 2, 102, 1002, 1102
		case 2:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = program[source1] * program[source2]
		case 102:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = source1 * program[source2]
		case 1002:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = program[source1] * source2
		case 1102:
			source1 = program[i+1]
			source2 = program[i+2]
			dest = program[i+3]
			step = 4
			program[dest] = source1 * source2
		case 3:
			dest = program[i+1]
			step = 2
			fmt.Println("Storing input of ", input, " at position ", dest)
			program[dest] = input
		case 4:
			dest = program[i+1]
			step = 2
			fmt.Println("Output at position ", i, " is ", program[dest])
		default:
			fmt.Println("Invalid opcode: ", opcode)
		}
	}

	return program[0]
}
