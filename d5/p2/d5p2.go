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

	run_program(program, 5)
}

func run_program(program []int, input int) int {
	var opcode, dest int
	max_index := len(program) - 1
	step := 1

	for i := 0; i <= max_index-3 && program[i] != 99; i += step {

		opcode = program[i]
		//		fmt.Println("Instruction at position ", i, ": ", opcode)
		switch opcode {
		// addition: 1, 101, 1001, 1101
		case 1, 101, 1001, 1101:
			dest = program[i+3]
			step = 4
			program[dest] = add(opcode, i, program)
		// multiplication: 2, 102, 1002, 1102
		case 2, 102, 1002, 1102:
			dest = program[i+3]
			step = 4
			program[dest] = multiply(opcode, i, program)
		// input
		case 3:
			dest = program[i+1]
			step = 2
			fmt.Println("Storing input of ", input, " at position ", dest)
			program[dest] = input
		// output
		case 4:
			dest = program[i+1]
			step = 2
			fmt.Println("Output at position ", i, " is ", program[dest])
		// jump if true
		case 5, 105, 1005, 1105:
			i, step = jump_if_true(opcode, i, program)
		// jump if false
		case 6, 106, 1006, 1106:
			i, step = jump_if_false(opcode, i, program)
		// less than
		case 7, 107, 1007, 1107:
			dest = program[i+3]
			step = 4
			program[dest] = less_than(opcode, i, program)
		// equals
		case 8, 108, 1008, 1108:
			dest = program[i+3]
			step = 4
			program[dest] = equals(opcode, i, program)
		default:
			fmt.Println("Invalid opcode: ", opcode)
		}
	}

	return program[0]
}

func add(op int, i int, program []int) int {
	source1 := program[i+1]
	source2 := program[i+2]

	switch op {
	case 1:
		return program[source1] + program[source2]
	case 101:
		return source1 + program[source2]
	case 1001:
		return program[source1] + source2
	case 1101:
		return source1 + source2
	}

	return -1
}

func multiply(op int, i int, program []int) int {
	source1 := program[i+1]
	source2 := program[i+2]

	switch op {
	case 2:
		return program[source1] * program[source2]
	case 102:
		return source1 * program[source2]
	case 1002:
		return program[source1] * source2
	case 1102:
		return source1 * source2
	}

	return -1
}

func jump_if_true(op int, i int, program []int) (int, int) {
	val := program[i+1]
	dest := program[i+2]
	switch op {
	case 5:
		if program[val] != 0 {
			return program[dest], 0
		}
	case 105:
		if val != 0 {
			return program[dest], 0
		}
	case 1005:
		if program[val] != 0 {
			return dest, 0
		}
	case 1105:
		if val != 0 {
			return dest, 0
		}
	}

	return i, 3
}

func jump_if_false(op int, i int, program []int) (int, int) {
	val := program[i+1]
	dest := program[i+2]
	switch op {
	case 6:
		if program[val] == 0 {
			return program[dest], 0
		}
	case 106:
		if val == 0 {
			return program[dest], 0
		}
	case 1006:
		if program[val] == 0 {
			return dest, 0
		}
	case 1106:
		if val == 0 {
			return dest, 0
		}
	}

	return i, 3
}

func less_than(op int, i int, program []int) int {
	source1 := program[i+1]
	source2 := program[i+2]

	switch op {
	case 7:
		if program[source1] < program[source2] {
			return 1
		}
	case 107:
		if source1 < program[source2] {
			return 1
		}
	case 1007:
		if program[source1] < source2 {
			return 1
		}
	case 1107:
		if source1 < source2 {
			return 1
		}
	}

	return 0
}

func equals(op int, i int, program []int) int {
	source1 := program[i+1]
	source2 := program[i+2]

	switch op {
	case 8:
		if program[source1] == program[source2] {
			return 1
		}
	case 108:
		if source1 == program[source2] {
			return 1
		}
	case 1008:
		if program[source1] == source2 {
			return 1
		}
	case 1108:
		if source1 == source2 {
			return 1
		}
	}

	return 0
}
