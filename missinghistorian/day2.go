/*-------------------------------------
Task1: 314      [OK]    Time: 1.1995ms
Task2: 373      [OK]    Time: 532.6µs
---------------------------------------*/

package missinghistorian

import (
	"math"
	"strconv"
	"strings"
)

type Rudolf struct {
	Solve
}

func (Rudolf *Rudolf) safeLine(i int, j int, data []int) bool {
	if data[i] < data[j] || math.Abs(float64(data[i] - data[j])) < 1 || math.Abs(float64(data[i] - data[j])) > 3 {
		return false
	} else if len(data) == 2 {
		return true
	} else {
		return Rudolf.safeLine(i, j, data[1:])
	}
}

func (Rudolf *Rudolf) filterLines(data string) string {
	var i_settings []int
	ret := 0
	rows := strings.Split(data, "\n")
	for _, r := range rows {
		if len(r) == 0 {
			continue
		}
		settings := strings.Split(r, " ")
		i_settings = []int{}
		for _, s := range settings {
			val, _ := strconv.Atoi(s)
			i_settings = append(i_settings, val)
		}
		if i_settings[0] > i_settings[1] && Rudolf.safeLine(0, 1, i_settings) {
			ret += 1
		} else if i_settings[0] < i_settings[1] && Rudolf.safeLine(1, 0, i_settings) {
			ret += 1
		}
	}
	return strconv.Itoa(ret)
}

func (Rudolf *Rudolf) filterErrorLines(data string) string {
	var i_settings []int
	
	ret := 0
	rows := strings.Split(data, "\n")
	for _, r := range rows {
		if len(r) == 0 {
			continue
		}
		settings := strings.Split(r, " ")
		i_settings = []int{}
		for _, s := range settings {
			val, _ := strconv.Atoi(s)
			i_settings = append(i_settings, val)
		}
		if i_settings[0] > i_settings[1] && Rudolf.safeLine(0, 1, i_settings) {
			ret += 1
		} else if i_settings[0] < i_settings[1] && Rudolf.safeLine(1, 0, i_settings) {
			ret += 1
		} else {
			for i := range i_settings {
				var levels = make([]int, len(i_settings))
				copy(levels, i_settings)
				if i == 0 {
					levels = levels[1:]
				} else if i < len(levels) {
					levels = append(levels[:i], levels[i+1:]...)
				} else {
					levels = levels[:len(levels)-1]
				}
				if levels[0] > levels[1] && Rudolf.safeLine(0, 1, levels) {
					ret += 1
					break
				} else if levels[0] < levels[1] && Rudolf.safeLine(1, 0, levels) {
					ret += 1
					break
				}
			}
		}
	}
	return strconv.Itoa(ret)
}

func (Rudolf *Rudolf) Task1(data string) string {
	return Rudolf.filterLines(data)
}

func (Rudolf *Rudolf) Task2(data string) string {
  return Rudolf.filterErrorLines(data)
}