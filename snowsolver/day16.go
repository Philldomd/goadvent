/*----------------------------------------
Task1: 7632     [OK]    Time: 1.6043ms
Task2: 8023     [OK]    Time: 201.7412ms
------------------------------------------*/

package snowsolver

import (
	"strconv"
	"strings"
)

type TheFloorWillBeLava struct {
	Solve
}

type LavaRock struct {
	pos Vector
	dir Vector
}

var LavaRock_List []LavaRock

func (theFloorWillBeLava TheFloorWillBeLava) Walk(mirror_map [][]rune, start_pos Vector, direction Vector, lava_map [][]rune) {
	for {
		next_pos := start_pos.Add(direction)
		if next_pos.x < 0 || next_pos.x >= len(mirror_map) {
			return
		}
		if next_pos.y < 0 || next_pos.y >= len(mirror_map[0]) {
			return
		}
		lava_map[next_pos.x][next_pos.y] = '#'
		switch mirror_map[next_pos.x][next_pos.y] {
		case '.':
			start_pos = start_pos.Add(direction)
		case '|':
			for _, save := range LavaRock_List {
				if save.pos == next_pos && save.dir == direction {
					return
				}
			}
			LavaRock_List = append(LavaRock_List, LavaRock{next_pos, direction})
			start_pos = start_pos.Add(direction)
			if direction.Equal(Vector{0, 1}) || direction.Equal(Vector{0, -1}) {
				theFloorWillBeLava.Walk(mirror_map, start_pos, Vector{1, 0}, lava_map)
				theFloorWillBeLava.Walk(mirror_map, start_pos, Vector{-1, 0}, lava_map)
				return
			}
		case '-':
			for _, save := range LavaRock_List {
				if save.pos == next_pos && save.dir == direction {
					return
				}
			}
			LavaRock_List = append(LavaRock_List, LavaRock{next_pos, direction})
			start_pos = start_pos.Add(direction)
			if direction.Equal(Vector{1, 0}) || direction.Equal(Vector{-1, 0}) {
				theFloorWillBeLava.Walk(mirror_map, start_pos, Vector{0, 1}, lava_map)
				theFloorWillBeLava.Walk(mirror_map, start_pos, Vector{0, -1}, lava_map)
				return
			}
		case '\\':
			for _, save := range LavaRock_List {
				if save.pos == next_pos && save.dir == direction {
					return
				}
			}
			LavaRock_List = append(LavaRock_List, LavaRock{next_pos, direction})
			start_pos = start_pos.Add(direction)
			temp := direction.x
			direction.x = direction.y
			direction.y = temp
		case '/':
			for _, save := range LavaRock_List {
				if save.pos == next_pos && save.dir == direction {
					return
				}
			}
			LavaRock_List = append(LavaRock_List, LavaRock{next_pos, direction})
			start_pos = start_pos.Add(direction)
			temp := -direction.x
			direction.x = -direction.y
			direction.y = temp
		}
	}
}

func (theFloorWillBeLava TheFloorWillBeLava) create_maps(data string) ([][]rune, [][]rune) {
	m_m := [][]rune{}
	l_m := [][]rune{}
	for _, row := range strings.Split(data, "\n") {
		if len(row) == 0 {
			continue
		}
		row_list := []rune{}
		lava_list := []rune{}
		for _, r := range row {
			row_list = append(row_list, r)
			lava_list = append(lava_list, '.')
		}
		m_m = append(m_m, row_list)
		l_m = append(l_m, lava_list)
	}

	return m_m, l_m
}

func (theFloorWillBeLava *TheFloorWillBeLava) Task1(data string) string {
	mirror_map, lava_map := theFloorWillBeLava.create_maps(data)
	theFloorWillBeLava.Walk(mirror_map, Vector{0, -1}, Vector{0, 1}, lava_map)
	sum := 0
	for _, lava := range lava_map {
		for _, l := range lava {
			if l == '#' {
				sum++
			}
		}
	}
	return strconv.Itoa(sum)
}

func (theFloorWillBeLava *TheFloorWillBeLava) Task2(data string) string {
	mirror_map, lava_map := theFloorWillBeLava.create_maps(data)
	R := len(mirror_map)
	C := len(mirror_map[0])
	sum_list := []int{}
	ret := 0
	for i := 0; i < C; i++ {
		sum := 0
		theFloorWillBeLava.Walk(mirror_map, Vector{-1, i}, Vector{1, 0}, lava_map)
		for l, lava := range lava_map {
			for m, la := range lava {
				if la == '#' {
					lava_map[l][m] = '.'
					sum++
				}
			}
		}
		LavaRock_List = []LavaRock{}
		sum_list = append(sum_list, sum)
		sum = 0
		theFloorWillBeLava.Walk(mirror_map, Vector{R, i}, Vector{-1, 0}, lava_map)
		for l, lava := range lava_map {
			for m, la := range lava {
				if la == '#' {
					lava_map[l][m] = '.'
					sum++
				}
			}
		}
		sum_list = append(sum_list, sum)
		LavaRock_List = []LavaRock{}
	}
	for i := 0; i < R; i++ {
		sum := 0
		theFloorWillBeLava.Walk(mirror_map, Vector{i, -1}, Vector{0, 1}, lava_map)
		for l, lava := range lava_map {
			for m, la := range lava {
				if la == '#' {
					lava_map[l][m] = '.'
					sum++
				}
			}
		}
		LavaRock_List = []LavaRock{}
		sum_list = append(sum_list, sum)
		sum = 0
		theFloorWillBeLava.Walk(mirror_map, Vector{i, C}, Vector{0, -1}, lava_map)
		for l, lava := range lava_map {
			for m, la := range lava {
				if la == '#' {
					lava_map[l][m] = '.'
					sum++
				}
			}
		}
		sum_list = append(sum_list, sum)
		LavaRock_List = []LavaRock{}
	}
	for _, sum := range sum_list {
		if ret < sum {
			ret = sum
		}
	}
	return strconv.Itoa(ret)
}
