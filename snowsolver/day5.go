/*---------------------------------------------
Task1: 382895070        [nil]   Time: 101.1µs
Task2: 17729182 [nil]   Time: 1h4m24.4294871s
-----------------------------------------------*/

package snowsolver

import (
	"slices"
	"strconv"
	"strings"
)

const MaxInt = int(9223372036854775807)

type Fertilizer struct {
	Solve
}

func (fertilizer *Fertilizer) Task1(data string) string {
	rows := strings.Split(data, "\n\n")
	seeds := []int{}
	for _, r := range rows {
		fields := strings.Split(r, ":")
		switch fields[0] {
		case "seeds":
			for _, seed := range strings.Fields(fields[1]) {
				korn, _ := strconv.Atoi(seed)
				seeds = append(seeds, korn)
			}
		default:
			fields[1] = fields[1][1:]
			for i, s := range seeds {
				for _, so := range strings.Split(fields[1], "\n") {
					if len(so) == 0 {
						continue
					}
					soil := strings.Fields(so)
					d_start, _ := strconv.Atoi(soil[0])
					s_start, _ := strconv.Atoi(soil[1])
					r_start, _ := strconv.Atoi(soil[2])
					if s >= s_start && s <= s_start+r_start-1 {
						new_seed := d_start + s - s_start
						seeds[i] = new_seed
					}
				}
			}
		}
	}
	return strconv.Itoa(slices.Min(seeds))
}

func (fertilizer *Fertilizer) Task2(data string) string {
	rows := strings.Split(data, "\n\n")
	seeds := []int{}
	seed_matrix := [][]int{}
	minmin := []int{}
	for _, seed := range strings.Fields(strings.Split(rows[0], ":")[1]) {
		korn, _ := strconv.Atoi(seed)
		seeds = append(seeds, korn)
	}
	rows = rows[1:]
	for _, r := range rows {
		temp_matrix := []int{}
		for _, m := range strings.Split(strings.Split(r, ":")[1], "\n") {
			matrix := strings.Fields(m)
			if len(matrix) == 0 {
				continue
			}
			n1, _ := strconv.Atoi(matrix[0])
			n2, _ := strconv.Atoi(matrix[1])
			n3, _ := strconv.Atoi(matrix[2])
			seed_matrix = append(seed_matrix, []int{n1, n2, n3})
		}
		seed_matrix = append(seed_matrix, temp_matrix)
	}
	for i := 0; i < len(seeds); i += 2 {
		min := MaxInt
		start := seeds[i]
		number := seeds[i+1]
		for j := start; j < start+number; j++ {
			min_temp := j
			found := false
			for _, m := range seed_matrix {
				if len(m) == 0 {
					found = false
					continue
				}
				if !found {
					if min_temp >= m[1] && min_temp <= m[1]+m[2]-1 {
						found = true
						min_temp = m[0] + min_temp - m[1]
					}
				}
			}
			if min_temp < min {
				min = min_temp
			}
		}
		minmin = append(minmin, min)
	}
	return strconv.Itoa(slices.Min(minmin))
}
