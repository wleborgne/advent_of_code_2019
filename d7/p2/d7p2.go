package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type amplifier struct {
	phase   int
	input   int
	program []int
}

func main() {
	var amp_chain [5]amplifier
	phase_orig := []int{5, 6, 7, 8, 9}
	original_program := read_program()
	prog_length := len(original_program)
	//max_val := 0
	//max_phase := phase_orig
	var amp_channels [5]chan int

	for p := make([]int, len(phase_orig)); p[0] < len(p); nextPerm(p) {
		phase_combo := getPerm(phase_orig, p)
		fmt.Println("Trying phase combo ", phase_combo)
		// initialize amplifiers with original program and assigned phase
		// also init new channels
		for i := 0; i < 5; i++ {
			amp_chain[i].program = make([]int, prog_length)
			amp_chain[i].phase = phase_combo[i]
			copy(amp_chain[i].program, original_program)
			amp_channels[i] = make(chan int)
		}

		// start amplifiers with assigned phase on input channel
		// they should block until the phase is put on the channel
		for i := 0; i < 5; i++ {
			//var next_input []int
			//next_input = append(next_input, amp_chain[i].phase)
			//next_input = append(next_input, prior_output)
			go run_program(amp_chain[i].program, amp_channels[i], amp_channels[(i+1)%5])
			amp_channels[i] <- amp_chain[i].phase
			//prior_output = current_output[0]
		}

		// all amps should be blocked at this time; kick off the process with
		// 0 sent to the first amp input channel
		amp_channels[0] <- 0
	}
}

// permutation functions from https://stackoverflow.com/a/30230552
func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

// INTCODE STUFF HERE
func read_program() []int {
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

	return program
}

func run_program(program []int, input chan int, output chan int) {
	var opcode, dest int
	max_index := len(program) - 1
	step := 1

	for i := 0; i <= max_index-3 && program[i] != 99; i += step {

		opcode = program[i]
		//		fmt.Println("Instruction at position ", i, ": ", opcode)
		switch opcode % 10 {
		// addition: 1, 101, 1001, 1101
		case 1:
			dest = program[i+3]
			step = 4
			program[dest] = add(opcode, i, program)
		// multiplication: 2, 102, 1002, 1102
		case 2:
			dest = program[i+3]
			step = 4
			program[dest] = multiply(opcode, i, program)
		// input
		case 3:
			dest = program[i+1]
			step = 2
			program[dest] = <-input
			fmt.Println("Stored input value ", program[dest], " in location ", dest)
		// output
		case 4:
			dest = program[i+1]
			step = 2
			output <- program[dest]
			fmt.Println("Output at position ", i, " is ", program[dest])
		// jump if true
		case 5:
			i, step = jump_if_true(opcode, i, program)
		// jump if false
		case 6:
			i, step = jump_if_false(opcode, i, program)
		// less than
		case 7:
			dest = program[i+3]
			step = 4
			program[dest] = less_than(opcode, i, program)
		// equals
		case 8:
			dest = program[i+3]
			step = 4
			program[dest] = equals(opcode, i, program)
		default:
			fmt.Println("Invalid opcode: ", opcode)
		}
	}
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
