/*----------------------------------------
Task1: 29130    [nil]    Time: 11.5189ms
Task2: 33438    [nil]    Time: 6.5947ms
------------------------------------------*/

package solvers

import (
	"strconv"
	"strings"
)

type PointOfIncidence struct {
	Solve
}

func (pointOfIncidence PointOfIncidence) FindMirror(mirrors_map [][]string, smudge int) int {
	investigate := []int{}
	n_rows := len(mirrors_map)
	n_elements := len(mirrors_map[0])
	for i := 0; i < n_rows-1; i++ {
		found := false
		err := 0
		for j := 0; j < n_elements; j++ {
			if mirrors_map[i][j] != mirrors_map[i+1][j] && err < smudge {
				err += 1
			} else if mirrors_map[i][j] != mirrors_map[i+1][j] && err >= smudge {
				found = false
				break
			} else {
				found = true
			}
		}
		if found {
			investigate = append(investigate, i)
		}
	}
	found := []int{}
	for inv, i := range investigate {
		j := i + 1
		equal := true
		err := 0
		if i == 0 {
			for n := 0; n < n_elements; n++ {
				if mirrors_map[i][n] != mirrors_map[j][n] && err < smudge {
					err += 1
				} else if mirrors_map[i][n] != mirrors_map[j][n] && err >= smudge {
					equal = false
					break
				}
			}
		} else {
			for i >= 0 && j < n_rows {
				for n := 0; n < n_elements; n++ {
					if mirrors_map[i][n] != mirrors_map[j][n] && err < smudge {
						err += 1
					} else if mirrors_map[i][n] != mirrors_map[j][n] && err >= smudge {
						equal = false
						break
					}
				}
				if !equal {
					break
				}
				i--
				j++
			}
		}
		if equal && err == smudge {
			found = append(found, investigate[inv]+1)
		}
	}
	if len(found) == 0 {
		return 0
	} else {
	  return found[0]
	}
}

func (pointOfIncidence PointOfIncidence) CollectData(mountain_list string) ([][]string, [][]string) {
	v_m, h_m := [][]string{}, [][]string{}
	map_list := strings.Split(mountain_list, "\n")
	for j := 0; j < len(map_list[0]); j++ {
		v_m = append(v_m, []string{})
	}
	for i := 0; i < len(map_list); i++ {
		for j := 0; j < len(map_list[0]); j++ {
			if len(map_list[i]) != 0 {
				v_m[j] = append(v_m[j], string([]rune(map_list[i])[j]))
			}
		}
	}
	for i, element := range strings.Split(mountain_list, "\n") {
		if len(element) == 0 {
			continue
		}
		h_m = append(h_m, []string{})
		for _, el := range element {
			h_m[i] = append(h_m[i], string(el))
		}
	}
	return v_m, h_m
}

func (pointOfIncidence *PointOfIncidence) Task1(data string) string {
	sum := 0
	for _, mountain_list := range strings.Split(data, "\n\n") {
		mirrors_col, mirrors_hor := pointOfIncidence.CollectData(mountain_list)
		value := pointOfIncidence.FindMirror(mirrors_col, 0)
		if value == 0 {
			value = pointOfIncidence.FindMirror(mirrors_hor, 0) * 100
		}
		sum += value
	}
	return strconv.Itoa(sum)
}

func (pointOfIncidence *PointOfIncidence) Task2(data string) string {
	sum := 0
	for _, mountain_list := range strings.Split(data, "\n\n") {
		mirrors_col, mirrors_hor := pointOfIncidence.CollectData(mountain_list)
		value := 0
		c_value := pointOfIncidence.FindMirror(mirrors_col, 1)
		h_value := pointOfIncidence.FindMirror(mirrors_hor, 1)
		if (c_value != 0 && c_value < h_value) || h_value == 0 {
			value = c_value
		} else if h_value != 0 {
			value = h_value * 100
		} else {
			value = 0
		}
		sum += value
	}
	return strconv.Itoa(sum)
}
