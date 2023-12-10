package solvers

import (
	"fmt"
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
	found := false
	nest := 0
	start_pos := 0
	start := Pipe{}
	n_r := len(pipes)
	r_l := len(pipes[0])
	first_direction := 0
	for i := 0; i < n_r; i++ {
		first_direction = 0
		for j := 0; j < r_l; j++ {
			found = false
			for _, p := range loop_pipes {
				if p.pos[0] == i && p.pos[1] == j {
					found = true
					if pipes[i][j] == '-' {
						break
					}
					if !ra {
						if first_direction == 0 && p.before[0] == p.after[0] {
							ra = true
							if p.before[0] == -1 {
								first_direction = 1
							} else {
								first_direction = 2
							}
						} else if first_direction == 0 && p.before[0] == 0 {
							first_direction = 1
							start_pos = p.after[0]
						} else if first_direction == 0 && p.after[0] == 0 {
							first_direction = 2
							start_pos = p.before[0]
						}
						start = p
						if pipes[i][j] == 'L' {
						  fmt.Println(first_direction, start_pos, start)
						}
						fmt.Print(":")
						ra = true
						break
					} else {
						if start.before[0] == p.after[0] && start.before[1] == p.after[1] || start.after[0] == p.before[0] && start.after[1] == p.before[1] {
							ra = false
							fmt.Print(string(pipes[i][j]), ";")
							break
						} else if first_direction == 1 && p.before[0] != 0 && p.before[0] != start_pos {
						  fmt.Print(string(pipes[i][j]), ";")
							ra = false
							break
						} else if first_direction == 2 && p.after[0] != 0 && p.after[0] != start_pos {
							fmt.Print(string(pipes[i][j]), ";")
							ra = false
							break
						} else if p.after[0] != 0 || p.before[0] != 0 {
							ra = false
							fmt.Print(string(pipes[i][j]), ";")
							first_direction = 0
							break
						}
					}
					/*if !ra {
						if p.before[0] != 0 {
							start_pos = p.before[0]
							ra = true
							break
						} else if p.after[0] != 0 {
							start_pos = p.after[0]
							ra = true
							break
						}
					} else {
						temp := 0
						if p.before[0] != 0 {
							temp = p.before[0]
						} else if p.after[0] != 0 {
							temp = p.after[0]
						}
						if temp != 0 && temp != start_pos {
							fmt.Print(string(pipes[i][j]), temp, start_pos)
							switch start_pos {
							case 1:
								if temp == -1 {
									ra = false
									break
								}
							case -1:
								if temp == 1 {
									ra = false
									break
								}
							}
						}
					}
					if ra {
						fmt.Println(start_pos)
					}
					if ra && p.after[0] != 0 {
						if start_pos != p.before[0] || start_pos != p.after[0] {
							ra = false
							fmt.Print(string(pipes[i][j]))
						}
					}*/
				}
			}
			if ra {
				fmt.Print(string(pipes[i][j]))
			}
			if ra && !found {
				fmt.Print(ra)
				nest += 1
			}
		}
		fmt.Println()
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
	fmt.Println(ret)
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
	fmt.Println(ret)
	return strconv.Itoa(pipeMaze.getNest(ret, pipes[:3]))
}
