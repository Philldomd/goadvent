/*-------------------------------------------
Task1: 674      [OK]    Time: 11.6001427s
Task2: 773      [OK]    Time: 6m26.7338187s
---------------------------------------------*/

package solvers

import (
	"strconv"
	"strings"
)

type ClumsyCrucible struct {
	Solve
}

func (clumsyCrucible ClumsyCrucible) createMaze(data string) [][]int {
	m := [][]int{}
	for _, row := range strings.Split(data, "\n") {
		if len(row) == 0 {
			continue
		}
		m_r := []int{}
		for _, r := range row {
			m_r = append(m_r, int(r-'0'))
		}
		m = append(m, m_r)
	}
	return m
}

type LavaNode struct {
	coord     Vector
	direction Vector
	heat_loss int
}

type LavaState struct {
	coord     Vector
	direction Vector
}

func (lv LavaNode) Pop() (Vector, Vector, int) {
	return lv.coord, lv.direction, lv.heat_loss
}

func (clumsyCrucible ClumsyCrucible) Dijkstras(maze [][]int, min int, max int) int {
	LegalMoves := make(map[Vector][]Vector)
	LegalMoves[Vector{0, 0}] = []Vector{{1, 0}, {0, 1}}
	LegalMoves[Vector{0, -1}] = []Vector{{1, 0}, {-1, 0}}
	LegalMoves[Vector{1, 0}] = []Vector{{0, -1}, {0, 1}}
	LegalMoves[Vector{0, 1}] = []Vector{{1, 0}, {-1, 0}}
	LegalMoves[Vector{-1, 0}] = []Vector{{0, -1}, {0, 1}}
	destination_coord := Vector{len(maze) - 1, len(maze[0]) - 1}
	lava_heap := []LavaNode{}
	lava_heap = append(lava_heap, LavaNode{Vector{0, 0}, Vector{0, 0}, 0})
	heat_map := make(map[Vector]int)
	heat_map[Vector{0, 0}] = 0
	visited := map[LavaState]struct{}{}
	for len(lava_heap) != 0 {
		temp_heat := lava_heap[0].heat_loss
		index := 0
		for i, lh := range lava_heap {
      if lh.heat_loss <  temp_heat {
				temp_heat = lh.heat_loss
				index = i
			}
		}
		coord, direction, heat_loss := lava_heap[index].Pop()
		state := LavaState{coord, direction}
		lava_heap = append(lava_heap[:index], lava_heap[index+1:]...)

		if coord.Equal(destination_coord) {
			return heat_loss
		}

		if _, ok := visited[state]; ok {
			continue
		}

		visited[state] = struct{}{}

		for _, new_direction := range LegalMoves[direction] {
			new_heat_loss := heat_loss
			for i := 1; i < max+1; i++ {
				new_coord := Vector{coord.x + i*new_direction.x, coord.y + i*new_direction.y}
				if new_coord.x < 0 || new_coord.y < 0 || new_coord.x > destination_coord.x || new_coord.y > destination_coord.y {
					continue
				}
				new_heat_loss = new_heat_loss + maze[new_coord.x][new_coord.y]
				if i >= min {
					if _, ok := heat_map[new_coord]; ok {
						if heat_map[new_coord] > new_heat_loss {
							heat_map[new_coord] = new_heat_loss
						}
					} else {
						heat_map[new_coord] = new_heat_loss
					}
					lava_heap = append(lava_heap, LavaNode{new_coord, new_direction, new_heat_loss})
				}
			}
		}
	}
	return -1
}

func (clumsyCrucible *ClumsyCrucible) Task1(data string) string {
	maze := clumsyCrucible.createMaze(data)
	return strconv.Itoa(clumsyCrucible.Dijkstras(maze, 1, 3))
}

func (clumsyCrucible *ClumsyCrucible) Task2(data string) string {
	maze := clumsyCrucible.createMaze(data)
	return strconv.Itoa(clumsyCrucible.Dijkstras(maze, 4, 10))
}
