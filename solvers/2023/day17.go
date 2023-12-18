/*------------------------------------
--------------------------------------*/

package solvers

import (
	"fmt"
	"strings"
)

type ClumsyCrucible struct {
	Solve
}

type AStarNode struct {
	parent   *AStarNode
	position Vector
	g        int
	h        int
	f        int
}

type DijkstrasNode struct {
	distance  int
	position  Vector
	direction Vector
	steps     int
}

type State struct {
	direction Vector
	steps     int
}

func (aStarNode AStarNode) Equal(a AStarNode) bool {
	return aStarNode.position == a.position
}

func (aStarNode *AStarNode) CalcH(goal Vector) {
	_x := aStarNode.position.x - goal.x
	_y := aStarNode.position.y - goal.y
	if _x < 0 {
		_x = -_x
	}
	if _y < 0 {
		_y = -_y
	}
	aStarNode.h = _x + _y
}

func (aStarLava *AStarNode) CalcF() {
	aStarLava.f = aStarLava.g + aStarLava.h
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

func (clumsyCrucible ClumsyCrucible) aStarLava(maze [][]int, start Vector, end Vector) []AStarNode {
	start_node := AStarNode{nil, start, maze[0][0], 0, maze[0][0]}
	start_node.CalcH(end)
	start_node.CalcF()
	direction := Vector{0, 0}
	old_direction := Vector{0, 0}
	straight := 0
	end_node := AStarNode{nil, end, maze[len(maze)-1][len(maze[len(maze)-1])-1], 0, 0}
	end_node.CalcH(end)
	end_node.CalcF()

	open_list := []AStarNode{}
	closed_list := []AStarNode{}
	open_list = append(open_list, start_node)

	for len(open_list) != 0 {
		current_node := open_list[0]
		current_index := 0
		for idx, node := range open_list {
			if node.f < current_node.f {
				current_node = node
				current_index = idx
			}
		}

		open_list = append(open_list[:current_index], open_list[current_index+1:]...)
		closed_list = append(closed_list, current_node)

		if current_node.parent != nil {
			direction = Vector{current_node.parent.position.x - current_node.position.x, current_node.parent.position.y - current_node.position.y}
		}

		if direction.Equal(old_direction) {
			straight += 1
		} else {
			straight = 0
		}

		old_direction = direction

		if current_node.Equal(end_node) {
			path := []AStarNode{}
			current := current_node
			for current.parent != nil {
				path = append(path, current)
				current = *current.parent
			}
			return path
		}

		children := []AStarNode{}

		for _, new_position := range []Vector{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			node_position := current_node.position.Add(new_position)
			if node_position.x > len(maze)-1 || node_position.x < 0 || node_position.y > len(maze[len(maze)-1])-1 || node_position.y < 0 {
				continue
			}

			new_node := AStarNode{&current_node, node_position, maze[node_position.x][node_position.y], 0, 0}

			children = append(children, new_node)
		}
		for i := range children {
			child_closed := false
			for _, closed_child := range closed_list {
				if children[i].Equal(closed_child) {
					child_closed = true
				}
			}
			if child_closed {
				continue
			}

			if straight == 3 {
				continue
			}

			children[i].CalcH(end)
			children[i].CalcF()

			for _, open_node := range open_list {
				if children[i].Equal(open_node) && children[i].g > open_node.g {
					continue
				}
			}
			open_list = append(open_list, children[i])
		}
	}
	return nil
}

func (clumsyCrucible ClumsyCrucible) check_step_function(cd DijkstrasNode, nd Vector) int {
	if nd.Equal(cd.direction) {
		return cd.steps + 1
	}
	if !nd.Equal(cd.direction) {
		return 1
	}
	if nd.x == -cd.direction.x && nd.y == -cd.direction.y {
		return 0
	}
	return 0
}

func (clumsyCrucible ClumsyCrucible) DijkstrasLava(maze [][]int, start Vector, max_step int) map[Vector]int {
	rows, cols := len(maze), len(maze[0])
	directions := []Vector{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	distances := make(map[Vector]int)
	for i, r := range maze {
		for j := range r {
			distances[Vector{i, j}] = 9999999
		}
	}
	queue := []DijkstrasNode{}
	queue = append(queue, DijkstrasNode{0, start, Vector{0, 0}, 1})
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			new_step := clumsyCrucible.check_step_function(current, dir)
			nx, ny := current.position.x+dir.x, current.position.y+dir.y
			if new_step == 0 || new_step == max_step {
				continue
			}
			if 0 <= nx && nx < rows && 0 <= ny && ny < cols {
				new_dist := current.distance + maze[nx][ny]
				if new_dist < distances[Vector{nx, ny}] {
					distances[Vector{nx, ny}] = new_dist
					queue = append(queue, DijkstrasNode{new_dist, Vector{nx, ny}, dir, new_step})
				}
			}
		}
	}
	return distances
}

func (clumsyCrucible *ClumsyCrucible) Task1(data string) string {
	maze := clumsyCrucible.createMaze(data)
	width, height := len(maze[0])-1, len(maze)-1
	path := clumsyCrucible.aStarLava(maze, Vector{0, 0}, Vector{width, height})
	//path := clumsyCrucible.DijkstrasLava(maze, Vector{0, 0}, 4)
	sum := 0
	//fmt.Println(path[Vector{height, width}])
	for _, p := range path {
		fmt.Println(p.position, p.g)
		sum += p.g
	}
	fmt.Println(sum)
	return "nil"
}

func (clumsyCrucible *ClumsyCrucible) Task2(data string) string {
	return "nil"
}
