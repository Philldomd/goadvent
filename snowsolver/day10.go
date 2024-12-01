
/*----------------------------------------
Task1: 6890     [nil]   Time: 5.233704s
Task2: 453      [nil]   Time: 5.4208927s
------------------------------------------*/

package solvers

import (
	"strconv"
	"strings"
)

type PipeMaze struct {
	Solve
}

type Pipe struct {
	pos    []int
	before []int
	after  []int
}

func (pipeMaze PipeMaze) ContainsPipe(i int, j int, loop []Pipe) bool {
	for _, p := range loop {
		if p.pos[0] == i && p.pos[1] == j {
			return true
		}
	}
	return false
}

func (pipeMaze PipeMaze) Walk(start []int, pipes [][]rune, direction []int) ([]Pipe, bool) {
	new_direction := []int{0, 0}
	ret := []Pipe{{[]int{start[0], start[1]}, []int{direction[0], direction[1]}, []int{0, 0}}}
	if start[0] < 0 || start[0] > len(pipes) || start[1] < 0 || start[1] > len(pipes[0]) {
		panic("Out of bounds")
	}
	switch pipes[start[0]][start[1]] {
	case 'S':
		return ret, true
	case 'J':
		if direction[0] == 1 {
			new_direction[0], new_direction[1] = 0, -1
		} else if direction[1] == 1 {
			new_direction[0], new_direction[1] = -1, 0
		} else {
			return ret, false
		}
	case '7':
		if direction[0] == -1 {
			new_direction[0], new_direction[1] = 0, -1
		} else if direction[1] == 1 {
			new_direction[0], new_direction[1] = 1, 0
		} else {
			return ret, false
		}
	case 'F':
		if direction[0] == -1 {
			new_direction[0], new_direction[1] = 0, 1
		} else if direction[1] == -1 {
			new_direction[0], new_direction[1] = 1, 0
		} else {
			return ret, false
		}
	case 'L':
		if direction[0] == 1 {
			new_direction[0], new_direction[1] = 0, 1
		} else if direction[1] == -1 {
			new_direction[0], new_direction[1] = -1, 0
		} else {
			return ret, false
		}
	case '-':
		if direction[0] != 0 {
			return ret, false
		} else {
			new_direction = direction
		}
	case '|':
		if direction[1] != 0 {
			return ret, false
		} else {
			new_direction = direction
		}
	default:
		return []Pipe{}, false
	}
	start[0], start[1] = start[0]+new_direction[0], start[1]+new_direction[1]
	ret[0].after[0], ret[0].after[1] = new_direction[0], new_direction[1]
	sum, found := pipeMaze.Walk(start, pipes, new_direction)
	ret = append(ret, sum...)
	if found {
		ret[len(ret)-1].after = direction
	}
	return ret, found
}

func (pipeMaze PipeMaze) getNest(loop_pipes []Pipe, pipes [][]rune) int {
	ra := false
	nest := 0
	n_r := len(pipes)
	r_l := len(pipes[0])
	for i := 0; i < n_r; i++ {
		ra = false
		for j := 0; j < r_l; j++ {
			if pipeMaze.ContainsPipe(i, j, loop_pipes) {
				char := pipes[i][j]
				if char == '|' {
					ra = !ra
				} else if char == 'L' || char == 'F' { //If the pipe turns to the search direction
					next := 0
					for _, dash := range pipes[i][j+1:] {
						next += 1
						if dash != '-' {
							break
						}
					}
					j += next
					next_char := pipes[i][j]
					if (char == 'L' && next_char == '7') || // check if we have an s bend or a u bend
						(char == 'F' && next_char == 'J') { // a u bend would not change ra to inside or outside
						ra = !ra
					}
				}
			} else if ra {
				nest += 1
			}
		}
	}
	return nest
}

func (pipeMaze *PipeMaze) Task1(data string) string {
	found := false
	pipes := [][]rune{}
	for _, row := range strings.Split(data, "\n") {
		if len(row) == 0 {
			continue
		}
		pipes = append(pipes, []rune(row))
	}
	start := []int{0, 0}
	for i, p := range pipes {
		for j, s := range p {
			if s == 'S' {
				start[0], start[1] = i, j
				break
			}
			if start[0] != 0 || start[1] != 0 {
				break
			}
		}
	}
	direction := []int{0, 1}
	ret, found := pipeMaze.Walk([]int{start[0] + direction[0], start[1] + direction[1]}, pipes, direction)
	if found {
		return strconv.Itoa(len(ret) / 2)
	} else {
		return "nil"
	}
}

func (pipeMaze *PipeMaze) Task2(data string) string {
	pipes := [][]rune{}
	for _, row := range strings.Split(data, "\n") {
		if len(row) == 0 {
			continue
		}
		pipes = append(pipes, []rune(row))
	}
	start := []int{0, 0}
	for i, p := range pipes {
		for j, s := range p {
			if s == 'S' {
				start[0], start[1] = i, j
				break
			}
			if start[0] != 0 || start[1] != 0 {
				break
			}
		}
	}
	direction := []int{0, 1}
	ret, _ := pipeMaze.Walk([]int{start[0] + direction[0], start[1] + direction[1]}, pipes, direction)
	return strconv.Itoa(pipeMaze.getNest(ret, pipes))
}
