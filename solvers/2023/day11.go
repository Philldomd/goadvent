/*----------------------------------------------
Task1: 9556896  [nil]   Time: 6.9925ms
Task2: 685038186836     [nil]   Time: 4.7085ms
------------------------------------------------*/

package solvers

import (
	"strconv"
	"strings"
)

type CosmicExpansion struct {
	Solve
}

type Vector struct {
	x int
	y int
}

func (V Vector) Add(v Vector) Vector {
	V.x += v.x
	V.y += v.y
	return V
}

func (V Vector) Equal(v Vector) bool {
	if V.x == v.x && V.y == v.y {
		return true
	}
	return false
}

type Space struct {
	char rune
	v    Vector
}

func (cosmicExpansion CosmicExpansion) ExpandGalaxies(pic [][]Space, addition int) []Vector {
	found := false
	if addition > 1 {
		addition = addition - 1
	}
	empty_space_x := []int{}
	empty_space_y := []int{}
	for i := 0; i < len(pic); i++ {
		found = false
		for _, c := range pic[i] {
			if c.char == '#' {
				found = true
				break
			}
		}
		if !found {
			empty_space_x = append(empty_space_x, i)
		}
	}
	for i := 0; i < len(pic[0]); i++ {
		found = false
		for j := 0; j < len(pic); j++ {
			if pic[j][i].char == '#' {
				found = true
				break
			}
		}
		if !found {
			empty_space_y = append(empty_space_y, i)
		}
	}
	galaxies := []Vector{}
	for _, s := range pic {
		for _, c := range s {
			if c.char == '#' {
				new_x := 0
				new_y := 0
				//calc new x
				for _, e_s_x := range empty_space_x {
					if e_s_x < c.v.x {
						new_x += addition
					}
				}
				//calc new y
				for _, e_s_y := range empty_space_y {
					if e_s_y < c.v.y {
						new_y += addition
					}
				}
				galaxies = append(galaxies, Vector{c.v.x + new_x, c.v.y + new_y})
			}
		}
	}
	return galaxies
}

func (cosmicExpansion CosmicExpansion) CreatePicture(data string) [][]Space {
	picture := [][]Space{}
	for i, r := range strings.Split(data, "\n") {
		if len(r) == 0 {
			continue
		}
		space := []Space{}
		for j, c := range r {
			space = append(space, Space{c, Vector{i, j}})
		}
		picture = append(picture, space)
	}
	return picture
}

func (cosmicExpansion CosmicExpansion) CalculateDistance(galaxies []Vector) []int {
	dist := []int{}
	for i, g := range galaxies {
		n_g := galaxies[i+1:]
		if i == len(galaxies)-1 {
			break
		}
		for _, ng := range n_g {
			X := ng.x - g.x
			if X < 0 {
				X = -X
			}
			Y := ng.y - g.y
			if Y < 0 {
				Y = -Y
			}
			dist = append(dist, X+Y)
		}
	}
	return dist
}

func (cosmicExpansion CosmicExpansion) Task1(data string) string {
	sum := 0
	distances := cosmicExpansion.CalculateDistance(cosmicExpansion.ExpandGalaxies(cosmicExpansion.CreatePicture(data), 1))
	for _, in := range distances {
		sum += in
	}
	return strconv.Itoa(sum)
}

func (cosmicExpansion CosmicExpansion) Task2(data string) string {
	sum := 0
	distances := cosmicExpansion.CalculateDistance(cosmicExpansion.ExpandGalaxies(cosmicExpansion.CreatePicture(data), 1000000))
	for _, in := range distances {
		sum += in
	}
	return strconv.Itoa(sum)
}
