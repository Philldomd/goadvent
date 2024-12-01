/*------------------------------------------------
Task1: 7260     [nil]    Time: 2.0348827s
Task2: 1909291258644    [nil]    Time: 902.9134ms
--------------------------------------------------*/
/* Note the T2 part of today was solved with much help from forums
i tried to brute force it but to no luck. Some suggested storing values
along the way and then i implemented a DP from an python implementation*/

package snowsolver

import (
	"strconv"
	"strings"
)

type HotSprings struct {
	Solve
}

var saved = make(map[string]int)

func (hotSprings HotSprings) valid(positions []rune, values []int) bool {
	sum := 0
	found := []int{}
	for _, c := range positions {
		if c == '.' {
			if sum > 0 {
				found = append(found, sum)
				sum = 0
			}
		}
		if c == '#' {
			sum += 1
		}
	}
	if sum > 0 {
		found = append(found, sum)
	}
	if len(found) == len(values) {
		for i, f := range found {
			if values[i] != f {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func (hotSprings HotSprings) combinationsT1(positions []rune, values []int, index int) int {	
	if index == len(positions) {
		if hotSprings.valid(positions, values) {
			return 1
		} else {
			return 0
		}
	}
	switch positions[index] {
	case '?':
		dots := []rune{}
		dots = append(dots, append(positions[:index], append([]rune{'.'}, positions[index+1:]...)...)...)
		cube := []rune{}
		cube = append(cube, append(positions[:index], append([]rune{'#'}, positions[index+1:]...)...)...)
		return (hotSprings.combinationsT1(dots, values, index) + hotSprings.combinationsT1(cube, values, index))
	default:
		return hotSprings.combinationsT1(positions, values, index+1)
	}
}

func (hotSprings HotSprings) combinationsT2(positions []rune, values []int) int {
	key := string(positions)
	for _, i := range values {
		key += strconv.Itoa(i) + ","
	}
	if v, ok := saved[key]; ok {
		return v
	}
	if len(positions) == 0 {
		if len(values) == 0 {
		  return 1
		} else {
			return 0
		}
	}
	ret := 0
	switch positions[0] {
	case '?':
		dots := []rune{}
		dots = append(dots, append([]rune{'.'}, positions[1:]...)...)
		cube := []rune{}
		cube = append(cube, append([]rune{'#'}, positions[1:]...)...)
		return (hotSprings.combinationsT2(dots, values) + hotSprings.combinationsT2(cube, values))
	case '#':
		if len(values) == 0 || len(positions) < values[0]{
			saved[key] = 0
			return 0
		}
		for _, p := range positions[:values[0]] {
			if p == '.' {
				saved[key] = 0
				return 0
			}
		}
		if len(values) > 1 {
			if len(positions) < values[0]+1 || positions[values[0]] == '#' {
				saved[key] = 0
				return 0
			}
      ret = hotSprings.combinationsT2(positions[values[0]+1:], values[1:])
		  saved[key] = ret
		  return ret
		} else {
		  ret = hotSprings.combinationsT2(positions[values[0]:], values[1:])
		  saved[key] = ret
		  return ret
		}
	case '.':
		ret = hotSprings.combinationsT2(positions[1:], values)
		saved[key] = ret
		return ret
	}
	return 0
}

func (hotSprings *HotSprings) Task1(data string) string {
	sum := 0
	rows := strings.Split(data, "\n")
	for _, r := range rows {
		if len(r) == 0 {
			break
		}
		fields := strings.Fields(r)
		positions := []rune(fields[0])
		values := []int{}
		for _, v := range strings.Split(fields[1], ",") {
			iv, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			values = append(values, iv)
		}
		sum += hotSprings.combinationsT1(positions, values, 0)
	}
	return strconv.Itoa(sum)
}

func (hotSprings *HotSprings) Task2(data string) string {
	sum := 0
	rows := strings.Split(data, "\n")
	for _, r := range rows {
		if len(r) == 0 {
			break
		}
		fields := strings.Fields(r)
		positions := []rune{}
		values := []int{}
		for i := 0; i < 5; i++ {
			if i < 4 {
				positions = append(positions, append([]rune(fields[0]), []rune{'?'}...)...)
			} else {
				positions = append(positions, []rune(fields[0])...)
			}
			for _, v := range strings.Split(fields[1], ",") {
				iv, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				values = append(values, iv)
			}
		}
		sum += hotSprings.combinationsT2(positions, values)
	}
	return strconv.Itoa(sum)
}
