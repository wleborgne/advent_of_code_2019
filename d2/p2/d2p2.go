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

	// mess with initial state to find expected output
	var noun, verb int
	var found = false

	for noun = 0; noun < 100 && !found; noun++ {
		for verb = 0; verb < 100 && !found; verb++ {
			fmt.Println("Trying noun: ", noun, ", verb: ", verb)
			new_program := make([]int, prog_length)
			copy(new_program, program)
			new_program[1] = noun
			new_program[2] = verb

			output := run_program(new_program)
			fmt.Println("Output: ", output)
			if output == 19690720 {
				fmt.Println("Found noun: ", noun, ", verb: ", verb, ", result: ", 100*noun+verb)
				found = true
			}
		}
	}
}

func run_program(program []int) int {
	var opcode, source1, source2, dest int
	max_index := len(program) - 1

	for i := 0; i <= max_index-3 && program[i] != 99; i += 4 {

		opcode = program[i]
		if program[i+1] <= max_index {
			source1 = program[i+1]
		} else {
			fmt.Println("Bad source1 index!!!")
			return 0
		}
		if program[i+2] <= max_index {
			source2 = program[i+2]
		} else {
			fmt.Println("Bad source2 index!!!")
			return 0
		}
		if program[i+3] <= max_index {
			dest = program[i+3]
		} else {
			fmt.Println("Bad dest index!!!")
			return 0
		}

		switch opcode {
		case 1:
			program[dest] = program[source1] + program[source2]
		case 2:
			program[dest] = program[source1] * program[source2]
		default:
			fmt.Println("Invalid opcode: ", opcode)
		}
	}

	return program[0]
}
